package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"parse"
)

func main(){
	// 获取屏幕输入的html文本内容来解析
	firstNode, err := html.Parse(os.Stdin)
	//firstNode, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v", err)
	}

	// 遍历打印节点
	parse.Visit(firstNode)
}