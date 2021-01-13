package parseImprove

import "golang.org/x/net/html"

// 该函数使用深度优先算法遍历一个node节点下的所有子节点和子节点的相邻节点（但不遍历自己的相邻节点）和定制化处理节点
// prev 和 next 是节点遍历前的操作和遍历后的操作
// ForeachNode只是遍历节点，真正对节点的操作在prev和next中
func ForeachNode(n *html.Node, prev, next func(*html.Node)){
	if prev != nil {
		prev(n)
	}

	// 递归遍历节点n的内部节点
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		ForeachNode(child, prev, next)
	}

	if next != nil {
		next(n)
	}
}