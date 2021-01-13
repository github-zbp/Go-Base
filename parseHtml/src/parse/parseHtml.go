package parse

import (
	"fmt"
	"golang.org/x/net/html"
)

/*
我们要通过html.parse(r *io.Reader)函数将一段html字符串转化为html节点对象，parse返回的是一个*html.Node的节点指针类型，而且返回的是html文本中的第一个节点。

func Parse(r io.Reader) (*Node, error)

所以我们可以先看一下这个html.Node节点类型到底包含些什么内容
type Node struct {
	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	Type      NodeType		// 节点类型
	DataAtom  atom.Atom
	Data      string		// 节点数据，如果是一个标签节点，那么这就是标签名
	Namespace string
	Attr      []Attribute	// 节点属性
}

parse返回的Node类型的变量包含了html中一个节点的真实内容（节点标签名，属性，标签类型），以及这个节点下的下一个节点的指针，子节点的指针，上一个节点的指针等，这些节点也都是Node类型。
由于parse返回的Node中只包含了一个节点A的真实内容，而其下的子节点和相邻节点的内容不存在A中，A只存他们的节点指针，所以parse返回的Node其实不大，很省内存空间。


我们可以再看看节点Node中的Type成员的类型 NodeType到底包含了哪些节点类型
type NodeType uint32
const (
	ErrorNode NodeType = iota		// 1
	TextNode						// 2	文本节点
	DocumentNode					// 3
	ElementNode						// 4	标签节点
	CommentNode						// 5
	DoctypeNode						// 6
	scopeMarkerNode					// 7
)

节点类型是一个32位的非负整型，一共有7中节点类型，被存到了常量下，

Attr成员是一个html.Attribute类型的切片，Attribute结构体保存这一个属性键值对。Attr []Attribute保存的就是一个节点的所有的属性键值对
type Attribute struct {
	Namespace, Key, Val string
}

 */

// 按照节点的结构遍历节点并输出, 接受一个要遍历的节点对象（html包中自定义的结构体）的指针
// 这里使用深度优先算法
func Visit(node *html.Node) {
	// 如果这个节点是标签节点才显示详细信息
	if node.Type == html.ElementNode {
		var nodeStr string = "<" + node.Data + " "

		//  遍历该节点所有的属性
		for _, attr := range node.Attr {
			keyValue := attr.Key + "=\"" + attr.Val + "\" "
			nodeStr += keyValue
		}

		nodeStr += ">"

		// 打印该节点
		fmt.Print(nodeStr + "\n")
	}

	// 如果这个节点有子节点(有子节点的不一定是标签节点，也可能是错误节点ErrorType或者文档节点DocumentType)，那么就递归子节点
	if node.FirstChild != nil {
		// 递归遍历打印该节点内部的节点
		Visit(node.FirstChild)
	}

	if node.Type == html.ElementNode {
		// 打印该节点的结束标签
		fmt.Printf("</%s>\n", node.Data)
	}

	// 如果这个节点有下一个节点，就递归下一个节点
	if node.NextSibling != nil {
		Visit(node.NextSibling)
	}

}

// 获取一个页面的所有a标签的链接
func GetLinks(node *html.Node, urlSlice []string) []string {
	// 判断是否为a标签
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr{
			if attr.Key == "href"{
				urlSlice = append(urlSlice, attr.Val)
			}
		}
	}

	// 判断是否有子节点，有则递归
	if node.FirstChild != nil {
		urlSlice = GetLinks(node.FirstChild, urlSlice)
	}

	// 判断是否有相邻节点,如果这个节点有下一个节点，就递归下一个节点
	if node.NextSibling != nil {
		urlSlice = GetLinks(node.NextSibling, urlSlice)
	}

	return urlSlice
}