package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var results []int

	if root == nil {
		return results
	}

	var tmpStack []*TreeNode

	tmpStack = append(tmpStack, root)

	for len(tmpStack) != 0 {
		cur := tmpStack[len(tmpStack)-1]

		results = append(results, cur.Val)
		tmpStack = tmpStack[:len(tmpStack)-1]

		if cur.Right != nil {
			tmpStack = append(tmpStack, cur.Right) // 因为出栈,这里需要先放右节点
		}
		if cur.Left != nil {
			tmpStack = append(tmpStack, cur.Left)
		}
	}

	return results
}
