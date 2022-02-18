package main

func canPartition(nums []int) bool {
	var max int
	numsLen := len(nums)
	if numsLen <= 1 {
		return false
	}

	dp := make([][]bool, numsLen)
	sum := 0
	for i := 0; i < numsLen; i++ {
		if max < nums[i] {
			max = nums[i]
		}
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	if max > target {
		return false
	}

	for i := 0; i < numsLen; i++ {
		dp[i] = make([]bool, target+1)
	}

	// 初始化dp, 需要将i=0，j=0填满已知结果，以便后续动态规划
	for i := 0; i < numsLen; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < numsLen; i++ {
		for j := 1; j <= target; j++ {
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[numsLen-1][target]
}
