package main

import (
	"fmt"
	"os"
	"spider"
	"sync"
	"time"
)

	type GoroutineCounter struct{
		number int
		cond *sync.Cond
	}

var crawl_tasks chan string
var limitSpeed chan struct{}
var seen map[string]bool
var seen_lock sync.Mutex
	var counter *GoroutineCounter

	func (gc *GoroutineCounter) Add(n int) {
//fmt.Println("before Add Lock")
		gc.cond.L.Lock()
//fmt.Println("Add Lock")
		gc.number += n
		gc.cond.L.Unlock()
//fmt.Println("Add unLock")
	}

	func (gc *GoroutineCounter) Done() {
		gc.cond.L.Lock()
//fmt.Println("Done Lock")
		gc.number--
		if gc.number <=0 {
			gc.cond.Signal()
		}
		gc.cond.L.Unlock()
//fmt.Println("Done unLock")
	}

	func (gc *GoroutineCounter) Wait(maxZeroTimes int) {
		time.Sleep(1 * time.Second)
		curZeroTimes := 0
		gc.cond.L.Lock()
//fmt.Println("Wait Lock")
		defer func () {
			gc.cond.L.Unlock()
//fmt.Println("Wait unLock")
		}()
		for {
//fmt.Println("Empty rotate")
			if gc.number > 0 {
				curZeroTimes = 0
//fmt.Println("before waiting")
				gc.cond.Wait()
			}else{
				if curZeroTimes >= maxZeroTimes{
					break
				}
				curZeroTimes++
				time.Sleep(3 * time.Second / 10)
			}
		}
	}

func main() {
	defer func (st time.Time) {
		fmt.Println(float64(time.Now().UnixNano() - st.UnixNano()) / 1e9)
	}(time.Now())

	startUrls := os.Args[1:]
	domain := startUrls[0]		// 必须以根域名为第一参，如 http://www.baidu.com/
	htmlDir := "./html"

	crawl_tasks = make(chan string, 1000)
	seen = map[string]bool{}
		counter = &GoroutineCounter{number: 0, cond: sync.NewCond(&sync.Mutex{}),}
	limitSpeed = make(chan struct{}, 20)

	go func() {
		for _, startUrl := range startUrls {
			// 将初始url假如到channel
//fmt.Println("Starturl")
			counter.Add(1)
			crawl_tasks <- startUrl
//fmt.Println("Put Starturl")
		}
	}()

	// 控制main goroutine的结束（如果没有这个goroutine的话，main goroutine在爬取完成后是会被crawl_tasks的接收一直阻塞的）
	go func(c *GoroutineCounter) {
		c.Wait(10)

		// 关闭 crawl_tasks 就可以终止crawl_tasks的接收了
		close(crawl_tasks)
	}(counter)

	// 开始爬取
	for url := range crawl_tasks {
//fmt.Println(url)
		go func(url string) {
			defer counter.Done()
			// 判断是否已经爬取过url
			seen_lock.Lock()
			crawled := seen[url]
			seen_lock.Unlock()
			if crawled {
				return
			}

			// 没有爬过
			seen_lock.Lock()
			seen[url] = true
			seen_lock.Unlock()

			limitSpeed<- struct{}{}
			cont,err := spider.Crawl(url)
			if err != nil {
				//fmt.Println(err)
			}
			<-limitSpeed

			// 解析url
			new_urls := spider.ParseUrl(cont,domain)

			// 生成新任务。生成新的goroutine做这件事是因为防止添加添加新任务时阻塞住爬虫goroutine
			go func(new_urls map[string]struct{}){
				counter.Add(len(new_urls))
				for new_url := range new_urls {
					crawl_tasks <- new_url
				}
			}(new_urls)

			// 下载cont内容
			spider.Download(cont, url, htmlDir)
		}(url)
	}
}
