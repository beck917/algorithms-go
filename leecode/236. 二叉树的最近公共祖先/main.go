package main

import "fmt"

func main() {
	a := &TreeNode{1, nil, nil}
	b := &TreeNode{2, nil, nil}
	c := &TreeNode{3, nil, nil}

	a.Left = b
	a.Right = c

	d := lowestCommonAncestor(a, b, c)
	fmt.Println(d.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LinkedList struct {
	Node *TreeNode
	Next *LinkedList
}

var foundp bool
var foundq bool

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath := &LinkedList{nil, nil}
	qPath := &LinkedList{nil, nil}

	dfs(root, p, q, pPath, qPath)

	var parentNode *TreeNode
	pPath = pPath.Next
	qPath = qPath.Next
	for pPath != nil && qPath != nil {
		if pPath.Node == qPath.Node {
			parentNode = pPath.Node
		}

		pPath = pPath.Next
		qPath = qPath.Next
	}

	// 刷题的时候注意全局变量需要重置
	foundp = false
	foundq = false

	return parentNode
}

func dfs(node, p, q *TreeNode, pPath, qPath *LinkedList) {
	if node == nil {
		return
	}

	if foundp != true {
		pPath.Next = &LinkedList{node, nil}
	}

	if p == node {
		foundp = true
	}

	if foundq != true {
		qPath.Next = &LinkedList{node, nil}
	}

	if q == node {
		foundq = true
	}

	if pPath != nil {
		pPath = pPath.Next
	}
	if qPath != nil {
		qPath = qPath.Next
	}

	if node.Left != nil {
		dfs(node.Left, p, q, pPath, qPath) // 这里应该传入.next，这里很容易处所，传入ppath
	}
	if node.Right != nil {
		dfs(node.Right, p, q, pPath, qPath)
	}
	return
}
