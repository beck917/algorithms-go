package main

import "fmt"

var results [][]int
var used map[int]bool

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

/**
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func permute(nums []int) [][]int {
	results = make([][]int, 0)
	used = make(map[int]bool)
	backtrace(nums, used, []int{})

	return results
}

func backtrace(nums []int, used map[int]bool, path []int) {
	if len(path) == len(nums) {
		p := make([]int, len(path))
		copy(p, path)
		results = append(results, p)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true

			backtrace(nums, used, path)

			used[i] = false
			path = path[:len(path)-1]
		}
	}
}
