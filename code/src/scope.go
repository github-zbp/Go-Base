package main

import "fmt"

//var x int
x := 10
//x := 10
func main() {
	//x := 100
	fmt.Println(x)
}

func test() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'	// A - a就是大写和小写之间的差距，所以对一个英文字符 +A-a就是变为大写的意思
			fmt.Printf("%c", x) // 如果直接Println(x)得到的是Ascii码，%c是格式化为单个字符
		}
	}
	fmt.Printf("\n")
	fmt.Println(x)
}