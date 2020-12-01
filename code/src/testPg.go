package main

import (
	"./pg"		// 要通过相对路径的方式引入，否则会被Go认为是引入系统包
	"fmt"
)

func main() {
	var y pg.Myint = 100	// 想要使用pg包的Myint类型必须写为 pg.Myint
	fmt.Println(y, pg.X)	// 想要使用pg包的X全局变量也必须写为 pg.X
	fmt.Println(y.Add(10))
	fmt.Println(pg.Calculate(y))
	//fmt.Println(y.sub(10))	// 报错：由于sub是首字母小写的，因此不可导出，所以在本文件中找不到sub方法
	//fmt.Println(pg.y)		// 报错：理由同上
}

