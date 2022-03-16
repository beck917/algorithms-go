package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var results []string

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, prevsum int) int {
	if root == nil {
		return 0
	}

	sum := prevsum*10 + root.Val

	leftsum := dfs(root.Left, sum)
	rightsum := dfs(root.Right, sum)

	return leftsum + rightsum
}
