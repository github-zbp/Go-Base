package main

import (
	"fetch"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"parseImprove"
)

var depth int		// 用于记录当前打印的标签的深度，按照这个深度来缩进这个标签

// 带缩进的遍历一个html文档的所有节点
func main(){
	url := "http://www.zbpblog.com"
	body, err := fetch.FetchBody(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	topNode, err := html.Parse(body)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	parseImprove.ForeachNode(topNode, printTagAddTab, printTagSubTab)
}

func printTagAddTab(n *html.Node){
	if n.Type != html.ElementNode {
		return
	}

	htmlStr := "<" + n.Data + " "

	for _, attr := range n.Attr {
		htmlStr += attr.Key + "=\"" + attr.Val + "\" "
	}

	htmlStr += ">"
	fmt.Printf("%*s%s\n", depth*4, "", htmlStr)
	depth++
}

func printTagSubTab(n *html.Node){
	if n.Type != html.ElementNode {
		return
	}
	depth--
	htmlStr := "</" + n.Data + ">"
	fmt.Printf("%*s%s\n", depth*4, "", htmlStr)

}