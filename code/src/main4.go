package main

import "fmt"

type B interface {
	test()
}

type BInstance struct{}
func (b BInstance) test(){
	// do something
}

type A struct{
	B
	otherAttrs []string
}



func main() {
	var b B
	b = BInstance{}
	a := A{b, []string{}}
	fmt.Println(a)
}
