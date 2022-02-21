package main

func change(amount int, coins []int) int {
	coinsLen := len(coins)
	dp := make([][]int, coinsLen)

	max := amount

	for k, _ := range coins {
		dp[k] = make([]int, max+1)
	}

	// init
	for i := 0; i <= max; i++ {
		if i%coins[0] == 0 {
			dp[0][i] = 1
		}
	}

	for k, _ := range coins {
		dp[k][0] = 1
	}

	for i := 1; i < coinsLen; i++ {
		v := coins[i]
		for j := 1; j <= max; j++ {
			if v > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-v]
			}
		}
	}
	return dp[coinsLen-1][max]
}
