package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func maxPathSum(root *TreeNode) int {
	max = math.MinInt64

	maxSum(root)

	return max
}

func maxSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := Max(maxSum(root.Left), 0)
	right := Max(maxSum(root.Right), 0)

	sum := left + right + root.Val
	if sum > max {
		max = sum
	}

	return root.Val + Max(left, right)
}

func Max(left, right int) int {
	if left >= right {
		return left
	}
	return right
}
