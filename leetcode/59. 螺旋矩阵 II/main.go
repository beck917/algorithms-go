package main

// 还是剥洋葱的思路
// https://leetcode-cn.com/problems/spiral-matrix-ii/solution/spiral-matrix-ii-mo-ni-fa-she-ding-bian-jie-qing-x/
// 通过模拟的算法,上右下左的顺序遍历
func generateMatrix(n int) [][]int {
	top, bottom, left, right := 0, n-1, 0, n-1
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	num := 0
	for num < n*n {
		for i := left; i <= right; i++ {
			num++
			matrix[top][i] = num
		}
		top++
		for i := top; i <= bottom; i++ {
			num++
			matrix[i][right] = num // 注意索引i和上面反的
		}
		right--

		for i := right; i >= left; i-- { // 注意这里是从大到小遍历
			num++
			matrix[bottom][i] = num
		}
		bottom--
		for i := bottom; i >= top; i-- {
			num++
			matrix[i][left] = num
		}
		left++
	}

	return matrix
}
