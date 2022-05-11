package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var stack []int

func kthSmallest(root *TreeNode, k int) int {
	stack = make([]int, 0)
	dfs(root, k)

	if len(stack) < k {
		return -1
	}

	return stack[len(stack)-1]
}

func dfs(root *TreeNode, k int) {
	if root == nil {
		return
	}

	dfs(root.Left, k)
	if len(stack) == k {
		return
	}
	stack = append(stack, root.Val)
	dfs(root.Right, k)

}
