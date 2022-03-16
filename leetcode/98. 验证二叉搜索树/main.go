package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	return dfs(root, math.MaxInt64, math.MinInt64)

}

func dfs(root *TreeNode, prevLeft, prevRight int) bool {
	if root == nil {
		return true
	}

	//fmt.Println(root.Val, prevLeft, prevRight)
	if root.Val >= prevLeft || root.Val <= prevRight {
		return false
	}

	return dfs(root.Left, root.Val, prevRight) && dfs(root.Right, prevLeft, root.Val)
}
