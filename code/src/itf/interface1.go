package itf

import (
	"fmt"
)

type Animal interface {
	Move()
}

type Person struct {
	Name string
	Age int
}

type Dog struct {
	Age int
}

func (person Person) Move() {
	fmt.Println("person move by two feet")
}

func (dog Dog) Move() {
	fmt.Println("person move by four feet")
}