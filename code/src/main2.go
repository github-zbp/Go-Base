package main

import "fmt"

func main(){
	str1 := "世界"
	firstAlpha, strLen := showStr(str1)

	fmt.Printf("字符串str1的内容为 '%s', 第一个字节为 %c, 长度为 %d", str1, firstAlpha, strLen)
}

func showStr(str string) (firstAlpha uint8, strLen int) {
	firstAlpha = str[0]
	strLen = len(str)
	return firstAlpha, strLen
}


