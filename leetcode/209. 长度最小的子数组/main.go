package main

import "math"

// target = 7, nums = [2,3,1,2,4,3]
// 通过滑动窗口, 当总和>7的时候,就移动start指针,然后再次对比是否>目标数据
func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt64

	start := 0
	end := 0
	sum := 0
	for end < len(nums) {
		sum += nums[end] // 每次计算总和

		for sum >= target {
			if end-start+1 < result {
				result = end - start + 1
			}
			sum -= nums[start] // 滑动窗口,就是减去上一个值
			start++            // 移动start指针,滑动start
		}

		end++
	}

	if result == math.MaxInt64 {
		result = 0
	}

	return result
}
