package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{1, 2, 3}, 3))
}

var results [][]int

func combinationSum(candidates []int, target int) [][]int {
	results = make([][]int, 0)
	backtrace(candidates, 0, []int{}, target)

	return results
}

func backtrace(nums []int, index int, path []int, target int) {

	if 0 == target {
		p := make([]int, len(path))
		copy(p, path)
		results = append(results, p)
	}

	if target < 0 {
		return
	}

	for i := index; i < len(nums); i++ {
		path = append(path, nums[i])

		backtrace(nums, i, path, target-nums[i])

		path = path[:len(path)-1]
	}
}
