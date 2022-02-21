package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{0, 0, 0}

	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	// 先进行排序
	sort.Ints(nums)
	//fmt.Println(nums)

	results := make([][]int, 0)

	// 三指针，左右指针从最左和最右，逼近，直到相遇
	n := len(nums)
	for i := 0; i < n; i++ {
		left := i + 1
		right := len(nums) - 1

		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			tmpNum := nums[i] + nums[left] + nums[right]
			if tmpNum < 0 {
				left++
			} else if tmpNum > 0 {
				right--
			} else {
				results = append(results, []int{nums[i], nums[left], nums[right]})

				// 去重
				for left+1 <= n-1 && nums[left] == nums[left+1] {
					left++
				}
				for right-1 >= 0 && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}

	return results
}
