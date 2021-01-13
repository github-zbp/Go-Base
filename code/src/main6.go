package main

import "fmt"

func main() {
	c := make(chan int, 5)
	for i:=0; i < 100; i++ {
		select {
			case x := <- c:
				fmt.Println(x)
			case c <- i:
		}
	}
}
