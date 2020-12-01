package main

import "fmt"

type Myint int

func main() {
	//var x Myint = 10
	//y := 10
	//fmt.Println(x)
	//fmt.Println(x == y)		// 报错，因为Myint和int不是相同类型，不能比较
	//y = x	// 也会报错，因为不能把给一个变量赋值一个不同类型的变量

	// example2
	//var x Myint = 10
	//y := 100
	//fmt.Println(x + Myint(y))

	// example3
	var x Myint = 1
	fmt.Println(x.Add(5))
}

func (x Myint) Add(n Myint) Myint {
	x += n
	return x
}