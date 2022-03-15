package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int

	var tmpStack []*TreeNode

	cur := root
	for cur != nil || len(tmpStack) > 0 {
		if cur != nil {
			tmpStack = append(tmpStack, cur)
			cur = cur.Left
		} else {
			cur = tmpStack[len(tmpStack)-1]
			result = append(result, cur.Val)
			tmpStack = tmpStack[:len(tmpStack)-1]
			cur = cur.Right
		}
	}

	return result
}
