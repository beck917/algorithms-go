package main

import (
	"math"
)

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	var mid int
	for left <= right {
		mid = (left + right) >> 1
		if get(nums, mid) > get(nums, mid+1) && get(nums, mid) > get(nums, mid-1) {
			return mid
		}

		if get(nums, mid-1) > get(nums, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return mid
}

func get(nums []int, index int) int {
	if index < 0 || index > len(nums)-1 {
		return math.MinInt64
	}

	return nums[index]
}
