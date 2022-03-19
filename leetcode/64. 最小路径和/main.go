package main

import "fmt"

func main() {
	grid := make([][]int, 1)
	//grid[0] = []int{1, 3, 1}
	//grid[1] = []int{1, 5, 1}
	//grid[2] = []int{4, 2, 1}
	grid[0] = []int{9, 1, 4, 8}

	fmt.Println(minPathSum(grid))
}

/**
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
*/
func minPathSum(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}

	dp := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
		if i == 0 {
			dp[0][0] = grid[0][0]
		}

		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			}

			if j == 0 && i != 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			}
		}
	}

	if len(grid) <= 1 {
		return dp[0][len(grid[0])-1]
	}

	i := 1
	j := 1
	for i = 1; i < len(grid); i++ {
		for j = 1; j < len(grid[i]); j++ {
			dp[i][j] = min(grid[i][j]+dp[i-1][j], grid[i][j]+dp[i][j-1])
		}
	}

	return dp[i-1][j-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
