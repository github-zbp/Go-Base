package main

import "fmt"

func main() {
	var str string
	str = "hello word 你好世界"
	fmt.Println(str)
}

func iter1(str string){
	length := len(str)
	for i := 0; i < length; i++ {
		fmt.Println(i, str[i])
	}
}

func iter2(str string){
	for i, chr := range str {
		fmt.Println(i, chr)
	}
}

func iter3(str string){
	for i, chr := range str {
		fmt.Println(i, string(chr))
	}
}