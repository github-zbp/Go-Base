package main

import "fmt"

func main() {
	//x := 2
	//p := &x		// 创建一个x的指针赋值给变量p
	//fmt.Println(p)	// p是x的地址，而不是值本身，是一串16进制的内容：0xc00009e058
	//fmt.Println(*p)		// *p才是x的值本身：2
	//
	//*p = 10		// 修改p指针指向的空间的值内容,会影响到x
	//fmt.Println(x, *p)		// 10  10

	//arr := [5]int{1,2,3,4,5}
	//p3 := &arr[2]		// 创建一个指向arr第3个元素的指针
	//*p3 = 100			// 通过指针修改arr第3个元素的值
	//fmt.Println(arr)	// [1,2,100,4,5]
	//
	//slice := arr[:]		// 创建一个切片
	//p4 := &slice[3]		// 允许对切片的某个元素取址，其本质还是对数组的元素取址
	//fmt.Println(p4, *p4)
	//
	//map1 := map[string]int{"a" : 1, "b" : 2}
	//pb := &map1["b"]	// 报错，因为go禁止对map的元素取址，原因后面介绍到map的时候会再说
	//fmt.Println(f() == f())

	//p1 := f()
	//p2 := f()
	//fmt.Println(*p1, *p2)
	//x := 1
	//fmt.Println(add(&x))	// 返回2
	//fmt.Println(x)			// x也受影响，变成2

	//p1 := new(int)
	//p2 := new(int)
	//fmt.Println(p1, *p1, p2, *p2, p1 == p2)
	//
	//p3 := new([5]int)
	//p4 := new([10]int)
	//fmt.Println(p3 == p4)
	//fmt.Println(*p3)

	//a := 10
	//b := "10"
	//fmt.Println(a==b)

	var (
		arr1 []int
		arr2 []int
	)
	fmt.Println(arr1 == arr2)

}

func add(p *int) int {
	*p++
	return *p
}

func f() *int{
	a := 2
	return &a
}