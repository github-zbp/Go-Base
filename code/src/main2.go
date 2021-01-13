package main

import (
	"fmt"
	"time"
)

func main(){
	naturals := make(chan int)
	squares := make(chan int)

	// 已知数据都是从channel的右边进左边出（单向）
	go counter(naturals)	// naturals在counter是右开左闭的
	go squarer(squares, naturals)	// naturals在squarer中是左开右闭，squares在squarer是右开左闭
	printer(squares)	// printer直接在main goroutine中跑，而不另开协程
}

// counter
func counter(in chan<-int){
	for i:=0;i <= 10;i++ {
		time.Sleep(3 * time.Second / 10)		// 限一下速度
		in<-i
	}
	close(in)
}

// squarer
func squarer(in chan<-int, out <-chan int){
	for i := range out{	// 不断接收counter发送的i
		in <- i * i
	}
	close(in)
}

// printer
func printer(in <-chan int){
	for res := range in{	// 不断接收squarer的结果
		fmt.Println(res)
	}
}