package main

import (
	"container/list"
	"fmt"
)

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) >> 1

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

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	index := search(nums, 5)

	fmt.Println(index)
}

type Graph struct {
	v int
	adj []*list.List
}

func (graph *Graph) DFS() {

}
