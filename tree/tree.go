package main

import "fmt"

type tree struct {
	root *node
}

type node struct {
	val   string
	left  *node
	right *node
}

func main() {
	root := &node{"A", nil, nil}
	t := &tree{}
	t.root = root

	root.left = &node{"B", nil, nil}
	root.right = &node{"G", nil, nil}

	root.left.left = &node{"C", nil, nil}
	root.left.right = &node{"D", nil, nil}

	root.left.right.left = &node{"E", nil, nil}
	root.left.right.right = &node{"F", nil, nil}

	root.right.right = &node{"H", nil, nil}

	preOrder(t.root)
	fmt.Println()
	middleOrder(t.root)
	fmt.Println()
	postOrder(t.root)
	fmt.Println("\nmaxDepth:", maxDepth(t.root))
	fmt.Println("\nminDepth:", minDepth(t.root))
}

func preOrder(n *node) {
	if n == nil {
		return
	}
	fmt.Print(n.val)
	preOrder(n.left)
	preOrder(n.right)
}

func middleOrder(n *node) {
	if n == nil {
		return
	}
	middleOrder(n.left)
	fmt.Print(n.val)
	middleOrder(n.right)
}

func postOrder(n *node) {
	if n == nil {
		return
	}
	postOrder(n.left)
	postOrder(n.right)
	fmt.Print(n.val)
}

// maxDepth
// 以后续遍历的方式进行
// 先遍历到最左边子树, 计算出最左边子树的层深肯定是1
// 然后进行回溯, 计算右边子树, 继续递归
// 以当前节点为根节点,对比左子树层深和右子树层深,将层深大的+1(当前节点层级),得出当前节点的层深
func maxDepth(n *node) int {
	if n == nil {
		return 0
	}
	left := maxDepth(n.left)
	right := maxDepth(n.right)
	fmt.Println()
	fmt.Print(n.val)
	fmt.Print(left)
	fmt.Print(right)

	if left > right {
		return left + 1
	}
	return right + 1
}

// minDepth
// 关键的理解点,在于对比,每个root节点的左右层深
func minDepth(n *node) int {
	if n == nil {
		return 0
	}
	left := minDepth(n.left)
	right := minDepth(n.right)

	if left > right {
		return right + 1
	}
	return left + 1
}

// levelOrder
func levelOrder(n node) {
	//var levelNodes []string
	nodes := []node{n}

	for len(nodes) > 0 {

	}
}
