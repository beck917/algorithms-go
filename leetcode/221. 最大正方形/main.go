package main

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))

	max := 0

	n := len(matrix)
	// 先将第一排横竖都设置为本身的值, 并标记max
	for i := 0; i < n; i++ {
		m := len(matrix[i])
		dp[i] = make([]int, m)
		dp[i][0] = int(matrix[i][0] - '0')

		if i == 0 {
			for j := 0; j < len(matrix[i]); j++ {
				dp[0][j] = int(matrix[0][j] - '0')

				if dp[0][j] > max {
					max = dp[0][j]
				}
			}
		}

		if dp[i][0] > max {
			max = dp[i][0]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')

			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1 // 公式,就是上下,和左上角的最小值+1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}

	return max * max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
