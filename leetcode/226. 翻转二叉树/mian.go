package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	dfs(root)

	return root
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left)
	dfs(root.Right)

	root.Left, root.Right = root.Right, root.Left

}
