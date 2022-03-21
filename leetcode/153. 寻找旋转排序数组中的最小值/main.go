package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 1, 2}))
}

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) >> 1
		//fmt.Println(left, right, mid, nums[left])

		if nums[left] > nums[mid] {

			right = mid

		} else {
			left = mid + 1
		}
	}

	return nums[left]
}

func findMin2(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) >> 1
		//fmt.Println(left, right, mid, nums[left])

		if nums[right] > nums[mid] { //寻找无序的那一半, 注意这里一定要用right判断,right判断有序后,无序的那一半0,1也适用

			right = mid

		} else {
			left = mid + 1
		}
	}

	return nums[left]
}
