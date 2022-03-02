package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 6, 1, 2}
	fmt.Println(search(nums, 2))
}

/** 2,4,5,6,7,0,1
输入：nums = [4,5,6,7,0,1,2], target = 6,7,0,1,2,4,5
输出：4
*/
func searchold(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) >> 1 // 这里是在循环内, 特别容易出错,需要记住要覆盖mid的值, 移位是1..
		//fmt.Println(left, right, mid)
		if nums[left] > nums[mid] { //记住是对比数据..不是索引
			if target >= nums[left] || target < nums[mid] {
				right = mid - 1
			} else if target > nums[mid] {
				left = mid + 1
			} else {
				return mid
			}
			continue
		}

		if nums[right] < nums[mid] {
			if target <= nums[right] || target > nums[mid] {
				left = mid + 1
			} else if target < nums[mid] {
				right = mid - 1
			} else {
				return mid
			}
			continue
		}

		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {

		mid := (left + right) >> 1
		//fmt.Println(left, right, mid)
		if target == nums[mid] {
			return mid
		}

		// 数组的规律, 只有两种情况,要么左边有序,要么右边有序
		// 我们只要对比了有序的部分, 数字是否存在其中, 否则数字肯定在无序的部分
		// 7 0 1 2 3 4
		// 4 5 6 7 0 1 2
		if nums[left] > nums[mid] {
			// 右边有序
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 左边有序
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
