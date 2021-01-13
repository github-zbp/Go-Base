package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取一天前的时间对象（获取了时间对象后，想获取格式化时间还是时间戳就都很容易了）
	now := time.Now()
	lastHour,_ := time.ParseDuration("-1h")		// 返回的yesterday是一个Duration类型，本质是一个int64类型的纳数，这里的lastHour应该是一个负数
	fmt.Println(lastHour, fmt.Sprintf("%d", lastHour))
	timeObj := now.Add(lastHour)	// Add需要传入一个Duration类型的纳秒，返回time.Time类型的时间对象
	printTime(timeObj)

	// ParseDuration支持的字符串有 “ns”, “us” (or “ns”), “ms”, “s”, “m”, “h”，也就是说最大支持小时这个维度，不支持天，如果像获取一天前的时间对象可以这样
	yesterdayObj := now.Add(-time.Hour * 24)		// time.Hour 是Duration类型的小时化为的纳秒数
	printTime(yesterdayObj)
}

func getTimeStamp() string {
	now := time.Now()

	// 这里注意 %2d 是输出两位整数，不足的位用空格填充； %02d 是输出两位整数，不足的位用0填充
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func printTime(t time.Time) {
	fmt.Println(t)
}