package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 8, 8, 8, 10}, 11))
}

/**
通过二分查找,获取数组的左边界和右边界
*/
func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		//fmt.Println(left, right, mid)

		if target > nums[mid] {
			left = mid + 1
		} else if target <= nums[mid] {
			right = mid - 1
		}
	}

	return left
}

/**
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
*/
func searchRange(nums []int, target int) []int {
	results := make([]int, 2)
	results[0] = -1
	results[1] = -1

	if len(nums) == 0 {
		return results
	}

	left := binarySearch(nums, target)

	if left > len(nums)-1 || nums[left] != target {
		return results
	}

	right := binarySearch(nums, target+1)

	results[0] = left
	results[1] = right - 1

	return results
}
