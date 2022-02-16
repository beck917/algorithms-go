package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 1

	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	return findQuick(nums, 0, len(nums)-1, len(nums)-k)
}

func partition(nums []int, start, end int) int {
	pivot := end
	j := start
	for i := start; i < end; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	nums[pivot], nums[j] = nums[j], nums[pivot]
	return j
}

func findQuick(nums []int, start, end, k int) int {
	//if start >= end {
	//	return 0
	//}

	index := partition(nums, start, end)
	//fmt.Println(index, k, nums[index])
	if index > k {
		return findQuick(nums, start, index-1, k)
	} else if index < k {
		return findQuick(nums, index+1, end, k)
	}

	return nums[index]
}
