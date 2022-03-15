package main

import "fmt"

func main() {
	fmt.Println(sortArray1([]int{2, 4, 5, 1, 3, 6}))
}

func sortArray1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)

	return nums
}

func partition(nums []int, start int, end int) int {
	pivot := nums[end]

	j := start
	for i := start; i <= end; i++ {
		if nums[i] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	nums[j], nums[end] = nums[end], nums[j]
	return j
}

func quicksort(nums []int, start int, end int) {
	if start >= end {
		return
	}

	index := partition(nums, start, end)
	quicksort(nums, 0, index-1)
	quicksort(nums, index+1, end)
}
