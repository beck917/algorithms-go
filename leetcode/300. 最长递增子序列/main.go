package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 4, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

// [10,9,2,5,3,7,4,101,18]
// [1 ,1,1,2,2,3,3]
func lengthOfLIS(nums []int) int {
	dp := make([]int, 0)

	for range nums {
		dp = append(dp, 1)
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	result := 0

	for _, v := range dp {
		if result < v {
			result = v
		}
	}
	return result
}
