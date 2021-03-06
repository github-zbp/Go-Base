package main

import (
	"fetch"
	"fmt"
	"golang.org/x/net/html"
	"parse"
)

func main() {
	url := "http://www.zbpblog.com"

	// 请求url，得到响应的Body成员
	body, err := fetch.FetchBody(url)

	if err != nil {
		fmt.Printf("%v", err)
	}

	// 将body传给解析html的函数,返回html的顶级节点
	topNode, err := html.Parse(body)

	if err != nil {
		fmt.Printf("%v", err)
	}

	// 打印所有标签
	parse.Visit(topNode)
}
