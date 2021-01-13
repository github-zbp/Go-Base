package method

import "fmt"

// 创建一个单向链表节点类
type IntList struct {
	Value int
	Point *IntList
}

// 计算这个单项链表的元素值总和
func (intList *IntList) Sum() int {
	if intList == nil {
		return 0
	}

	return intList.Value + intList.Point.Sum()
}

// 往单项链表中添加一个元素，会影响原链表
func (intList *IntList) Add(i int) {
	// 获取链表最后一个节点的指针
	lastNode := intList
	for lastNode.Point != nil {
		lastNode = lastNode.Point
	}

	lastNode.Point = &IntList{
		Value: i,
		Point: nil,
	}
}

// 打印这个链表
func (IntList *IntList) IterPrint() {
	currentNode := IntList
	for currentNode != nil {
		fmt.Printf("%d\t", currentNode.Value)
		currentNode = currentNode.Point
	}
}