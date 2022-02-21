package main

import (
	"fmt"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

type nodePath struct {
	val  int
	next *nodePath
}

func (np *nodePath) sum() int {
	sum := 0
	for np != nil {
		sum += np.val
		fmt.Println("test", np.val)
		np = np.next
	}

	return sum
}

func main() {
	tree := &treeNode{1, nil, nil}
	tree.left = &treeNode{2, nil, nil}
	tree.right = &treeNode{3, nil, nil}
	tree.left.left = &treeNode{4, nil, nil}
	tree.left.right = &treeNode{5, nil, nil}

	np := &nodePath{0, nil}
	fmt.Println(pathSum(tree, np, np, 4))
}

func pathSum(node *treeNode, np, npnext *nodePath, targetSum int) bool {
	if node == nil {
		return false
	}

	npnext.next = &nodePath{node.val, nil}
	fmt.Println(npnext.next.val)
	//sum
	if np.sum() == targetSum {
		return true
	}

	return pathSum(node.left, np, npnext.next, targetSum) || pathSum(node.right, np, npnext.next, targetSum)
}
