package main

import "fmt"

func main() {
	array1 := [5] int {1,2,3,4,5}
	modify(array1)
	fmt.Println(array1)
}

func modify(array [5] int) {
	array[0] = 10
	fmt.Println(array)
}