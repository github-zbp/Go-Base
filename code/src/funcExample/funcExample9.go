package funcExample

import (
	"log"
	"time"
)

func Run9() {
	defer logTime("Run9")()

	time.Sleep(3 * time.Second)		// Sleep需要传入一个Duration类型，Duration是一个纳秒
}

func logTime(funcName string) func(){
	// 记录开始运行时间
	startTime := time.Now()
	log.Printf("函数 %s 开始运行", funcName)
	return func() {
		log.Printf("函数 %s 运行结束， 耗时 %s", funcName, time.Since(startTime))		// Since需要传入一个Time对象，返回Duration类型
	}
}