package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(nums))
}

/**
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

这是一个dp题, 做起来还是比较烦的

首先看一下普通解法,n方的解法,两次循环,找到最大的和就可以

但这题n2很容易超时,所以需要看一下dp的解法

   -2 1 3 4 5 -1 -12 5
dp -2 1 4 7 12 11 -1 5

这里的重要的公式推到,是dp[i-1]+nums[i] > nums[i]也就是前一个dp的数据和当前数据相加
看是否>当前数据,那么直接继续用当前数据

也很好理解,如果>当前数据,后面的累加还有可能超过,如果<当前数据,那么当前数据和后面的累加才是最大的,放弃前面的数据
*/
func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	max := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
