package main

import (
	"os"
	tt "text/template"
	"time"
)

// 定义书架（有多本书）结构体类型
type Books struct {
	TotalCount	int
	Items []Book	// 书架有多本书，里面存放的元素都是Book类型的结构体
}

// 定义书的结构体类型
type Book struct {
	Title string
	Author *Writer		// 结构体成员不能是结构体的类型，而只能是结构体的指针的类型，所以这里不能写为 Author Writer。用指针一样可以操控结构体
	CreateAt time.Time	// 书的创建时间是一个time.Time类型
}

// 定义作者的结构体类型
type Writer struct {
	Name string
	Country string
}

func main() {
	// 第一步 定义模板内容 用··定义多行字符串
	const templ = `{{.TotalCount}} issues:
{{range .Items}}
Author: {{.Author.Name}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

	// 这个模板内容里面 {{}} 是变量， {{range xxx}}是循环， {{xxx | xxx}} 是使用函数，包括go已有的函数和自定义函数。如果是自定义函数则需要将这个函数注册进模板对象中

	// 第二步 生成模板对象
	myTemplateFuncs := tt.FuncMap{"daysAgo" : daysAgo}		// FuncMap 是template包中自定义的哈希表类型
	//myTemplate, err := tt.New("myTemplate").Funcs(myTemplateFuncs).Parse(templ)		// 这里创建了一个template对象，并注册了自定义函数，并指定了要使用的模板内容
	//if err != nil {
	//	log.Fatal(err)
	//}
	myTemplate:= tt.Must(tt.New("myTemplate").Funcs(myTemplateFuncs).Parse(templ))

	// 第三步 往模板中注入变量，使用模板
	shengjingWriter := Writer{
		Name:    "Kernighan",
		Country: "unknown",
	}
	shengjingCreateAt, _ := time.Parse("2006-01","2016-03")
	shengjing := Book {		// go语言圣经
		Title:"The Go Programming Language",
		Author:&shengjingWriter,
		CreateAt:shengjingCreateAt,
	}

	var (
		concurrenyInGo Book
		concurrenyInGoWriter Writer
	)
	concurrenyInGoCreateAt, _ := time.Parse("2006-01", "2018-11")
	concurrenyInGo = Book {
		Title:"Go语言并发之道",
		Author:&concurrenyInGoWriter,
		CreateAt: concurrenyInGoCreateAt,
	}
	concurrenyInGoWriter = Writer {
		Name: "Katherine Cox-Buday",
		Country: "unknown",
	}

	books := Books {
		TotalCount: 2,
		Items: []Book{concurrenyInGo, shengjing},
	}

	// os.Stdout是标准输出，意思是将渲染好的模板内容直接打印到屏幕的意思
	myTemplate.Execute(os.Stdout, books)
	//if err := myTemplate.Execute(os.Stdout, books); err != nil {
	//	log.Fatal(err)
	//}
}


func daysAgo(createTime time.Time) int {
	// time.Since(t) 返回当前时间的时间对象与t这个时间对象之间相差的纳秒的绝对值，返回的是Duration类型的纳秒
	// time.Hour 是Duration类型的1个小时的纳秒，即 3600 * 10^9
	return int(time.Since(createTime) / time.Hour / 24)
}