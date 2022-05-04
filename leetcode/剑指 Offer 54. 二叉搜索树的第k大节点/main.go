package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nums []int

func kthLargest(root *TreeNode, k int) int {
	nums = make([]int, 0)
	dfs(root)

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1-k {
			return nums[i]
		}
	}

	return -1
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left)
	nums = append(nums, root.Val)
	dfs(root.Right)
}
