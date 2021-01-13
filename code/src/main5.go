package main

import "fmt"

var deposit int
func main() {
	waiting := make(chan struct{})
	go func(){
		for i:=0; i < 100000000; i++{
			deposit++
		}
		waiting <- struct{}{}
	}()

	for i:=0; i < 100000000; i++{
		deposit--
	}

	<-waiting
	fmt.Println(deposit)
}