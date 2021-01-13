package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main(){
	var w io.Writer
	w = os.Stdout

	w_methods := getAllMethod(w)
	fmt.Printf("%#v", w_methods)
}

func getAllMethod(x interface{}) (methods map[string]string){
	methods = map[string]string{}
	x_val := reflect.ValueOf(x)
	x_type := x_val.Type()

	// 遍历x的所有方法
	for i := 0; i < x_val.NumMethod(); i++{
		x_method := x_type.Method(i)
		methods[x_method.Name] = x_method.Type.String()
	}

	return methods
}