package main

func longestConsecutive(nums []int) int {
	n := len(nums)
	visited := make(map[int]int, n)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		if visited[nums[i]-1] != 0 {
			dp[i] += visited[nums[i]-1]
		}

		if visited[nums[i]+1] != 0 {
			dp[i] += visited[nums[i]+1]
		}

		visited[nums[i]] = dp[i]
	}

	return dp[n-1]
}
