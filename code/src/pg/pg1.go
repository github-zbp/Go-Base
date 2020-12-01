package pg

import "fmt"

type Myint int		// 在全局声明了一个自定义Myint类
var X Myint = 0		// 在全局声明了可导出的变量X

func init(){	// 定义包的初始化函数
	X++
	fmt.Println("pg1")
}

func (x Myint) Add(n Myint) Myint {		// 在包中定义一个可导出的类方法Add
	x += n
	return x
}

func PrintA() {
	fmt.Println(A)
}