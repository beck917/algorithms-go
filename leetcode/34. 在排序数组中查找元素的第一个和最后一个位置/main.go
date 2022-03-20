package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 8, 8, 8, 10}, 8))
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

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		//fmt.Println(left, right, mid)

		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid - 1
			right = mid + 1
			results[0] = mid
			results[1] = mid

			for left >= 0 || right < len(nums) {
				modify := false
				if left >= 0 && nums[left] == target {
					results[0] = left
					left--
					modify = true
				}

				if right < len(nums) && nums[right] == target {
					results[1] = right
					right++
					modify = true
				}

				if modify == false {
					break
				}
			}
			break
		}
	}

	return results
}
