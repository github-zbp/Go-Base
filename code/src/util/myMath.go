package util

import "fmt"

func MyAdd(a, b int) (c int)  {
	c = a + b
	return
}

func init() {
	fmt.Println("before ")
}