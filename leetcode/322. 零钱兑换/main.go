package main

import (
	"fmt"
)

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}

/**
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
dp[i] = dp[i][j-coin[i]]+1
*/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	//sort.Ints(coins)
	n := len(coins)
	dp := make([][]int, len(coins))

	for i := 0; i < n; i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j < amount+1; j++ {
			dp[i][j] = amount + 1 // 设置为最大值, 表示无解 这里不能用max.maxint64.dp[i][j-coins[i]]+1计算的时候会越界....
		}
	}

	dp[0][0] = 0
	for j := 1; j < amount+1; j++ {
		if j%coins[0] == 0 {
			dp[0][j] = j / coins[0]
		}
	}

	for i := 0; i < n; i++ {
		dp[i][0] = dp[0][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < amount+1; j++ {
			if j < coins[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i]]+1)
			}
			// fmt.Println(dp[i][j], i, j)

		}
	}

	if dp[n-1][amount] == amount+1 {
		return -1
	}

	return dp[n-1][amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
