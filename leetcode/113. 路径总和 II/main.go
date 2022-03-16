package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var results [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return results
	}

	results = make([][]int, 0)
	dfs(root, targetSum, []int{}, 0)

	return results
}

func dfs(root *TreeNode, targetSum int, path []int, sum int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	sum += root.Val

	if root.Left == nil && root.Right == nil {
		//fmt.Println(sum, targetSum)
		if sum == targetSum {
			p := make([]int, len(path))
			copy(p, path)
			results = append(results, p)
		}
		return
	}

	dfs(root.Left, targetSum, path, sum)
	dfs(root.Right, targetSum, path, sum)
}
