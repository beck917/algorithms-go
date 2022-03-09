package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	var dp []int
	for i := 0; i < n; i++ {
		dp = append(dp, 0)
	}

	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs2(n-1) + climbStairs2(n-2)
}
