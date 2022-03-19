package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

var results [][]int

/**
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
*/
func subsets(nums []int) [][]int {
	results = make([][]int, 0)
	results = append(results, []int{})
	backtrace(nums, 0, []int{})

	return results
}

func backtrace(nums []int, index int, path []int) {
	if index > len(nums)-1 {
		//fmt.Println(path)
		return
	}

	for i := index; i < len(nums); i++ {
		path = append(path, nums[i])
		p := make([]int, len(path))
		copy(p, path)
		results = append(results, p)
		backtrace(nums, i+1, path)
		path = path[:len(path)-1]
	}

	return
}
