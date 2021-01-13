package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var curSize, curFileNum int64
	var wg sync.WaitGroup		// 用一个WaitGroup来记录和等待所有的walkDir goroutine结束
	dirPath := os.Args[1]
	size := make(chan int64)
	limitSpeed := make(chan struct{}, 100)
	ticker := time.Tick(2 * time.Second)

	done := make(chan struct{})			// 用于接收用户终止信号的channel

	wg.Add(1)
	go walkDir(dirPath, size, limitSpeed, done, &wg)
	go func(wg *sync.WaitGroup) {		// 开了一个新协程等待walkDir协程结束，Wait被唤醒后关闭size这个channel从而解除main的阻塞
		wg.Wait()
		close(size)
	}(&wg)

	go func(){		// 开一个goroutine接收用户的终止指令
		os.Stdin.Read(make([]byte, 1))		// 会阻塞该协程
		close(done)		// 通过关闭done的方式广播给所有goroutine中的cancelled方法
	}()

loop:
	for {
		select {
		case s, ok := <-size: // 99.999%会走到这个case，因为walkDir一直在往size发送数据，所以size一直都可接收
			if !ok {
				break loop
			}
			curFileNum += 1
			curSize += s
		case <-ticker: // 每隔两秒会走到这个case
			fmt.Printf("当前已遍历 %d 个文件，大小为 %dKb \n", curFileNum, curSize/1000)
		case <-done: // 只有当用户在屏幕敲了东西，然后程序执行了close(done)才会走到这个case
			for range size {}		// 当用户停止程序，就从size中取出所有的数据但不作为；当所有的size都取出来之后，这里会阻塞，但是不要紧，很快所有的goroutine都退出，然后close(size)就会执行，然后这个for就会跳出执行到下面的break loop
			break loop
		}
	}


	fmt.Println(curFileNum, curSize)
}

// 判断用户是否发出终止指令的方法。cancelled是一个被动方法，他不会主动检测用户是否发出终止指令，而是在程序运行到cancelled的时候才会检测用户是否发出终止
// 运行到cancelled时，如果用户已经发出了终止指令(关闭了done这个channel)，那么done就会接收到这个指令（不阻塞的接收到一个零值），返回true；如果运行到cancelled时，如果用户没有发出终止指令也不会被done阻塞而是走的default，直接返回
func cancelled(done <-chan struct{}) bool {
	select {
		case <-done:
			return true
		default:
			return false
	}
}

func walkDir(dirPath string, size chan int64, limitSpeed, done chan struct{}, wg *sync.WaitGroup) (err error){
	defer wg.Done()

	// 判断是否已经取消
	//if cancelled(done){
	//	return nil
	//}

	limitSpeed <- struct{}{}
	// 判断是否已经取消
	if cancelled(done){
		<- limitSpeed
		return nil
	}
	dirPath = strings.TrimRight(dirPath, "/") + "/"
	fileInfos, err := ioutil.ReadDir(dirPath)
	<- limitSpeed
	if err != nil{
		fmt.Println(err)
		return err
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			wg.Add(1)
			go walkDir(dirPath + fileInfo.Name(), size, limitSpeed, done, wg)
		}else{
			size <- fileInfo.Size()
		}
	}

	return err
}
