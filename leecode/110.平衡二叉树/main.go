package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	level := getHeight(root)
	if level == -1 {
		return false
	}

	return true
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := getHeight(node.Left)
	if left == -1 {
		return -1
	}
	right := getHeight(node.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left-right)) > 1 {
		return -1
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
