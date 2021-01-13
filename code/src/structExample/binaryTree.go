package structExample

type BinaryTree struct {
	value int		// 二叉树根节点的值
	left *BinaryTree	// 二叉树根节点的左分支，指向的是一个左子树
	right *BinaryTree 	// 二叉树根节点的右分支，指向的是一个右子树
}

// 给一个切片排序（将切片的元素添加到一个二叉树中）,返回一个排好序的切片
// 是个导出函数
func Sort(values []int) []int{
	// 创建一个空的二叉树,tree是一个空指针，所以tree为nil（因为所有类型指针的零值都是nil）
	var tree *BinaryTree

	// 先将切片的元素存放到二叉树中，存放的过程就是排序的过程
	for _, value := range values {
		tree = add(value, tree)	 // 传入二叉树的指针, 这里必须返回tree，因为add外的tree和add内的tree是两个不同的指针了，我之前说过go的函数传参是值拷贝，所以如果不返回一个tree指针，add外的tree就还是一个nil空指针，add返回的tree才是一个指向构建好了的二叉树的指针
	}

	// 再将二叉树中的元素释放到一个新的切片中，这个新切片中的元素就是排好序的元素
	sortedSlice := appendValues(tree, []int{})

	return sortedSlice
}

// 往二叉树添加一个元素，add必须返回一个tree
// 是个非导出函数
func add(value int, tree *BinaryTree) *BinaryTree {
	if tree == nil {	// 如果这棵树没有初始化则初始化
		tree = new(BinaryTree)		// 在底层创建了一个匿名变量二叉树，并返回二叉树的指针
		tree.value = value
		return tree
	}

	if value > tree.value {
		tree.right = add(value, tree.right)			// 如果树的根节点有值了，而且插入的值比树节点的值大，则将这个值添加到树的右子树中
	}

	if value <= tree.value {
		tree.left = add(value, tree.left)			// 同上
	}
	return tree
}

// 将二叉树中的元素释放都一个切片中，得到的切片其里面的元素就已经排好序的元素
// 是个非导出函数
// values参数是一个空切片，用于接收二叉树释放的元素
func appendValues(tree *BinaryTree, values []int) []int {
	if tree == nil{
		return values
	}

	// 使用树的前序遍历来遍历树的所有元素
	values = appendValues(tree.left, values)
	values = append(values, tree.value)
	values = appendValues(tree.right, values)

	return values
}