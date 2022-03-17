package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	dfs(root)
	return maxDiameter - 1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	max := 0
	if left >= right {
		max = left + 1
	} else {
		max = right + 1
	}

	if left+right+1 > maxDiameter {
		maxDiameter = left + right + 1
	}

	return max

}
