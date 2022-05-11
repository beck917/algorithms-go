package main

import "sort"

/**
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
*/
var results [][]int

func permuteUnique(nums []int) [][]int {
	results = make([][]int, 0)
	visited := make(map[int]bool, 0)

	sort.Ints(nums)

	backtrace(nums, []int{}, visited)

	return results
}

func backtrace(nums []int, path []int, visited map[int]bool) {
	if len(path) == len(nums) {
		p := make([]int, len(path))
		copy(p, path)
		results = append(results, p)
		return
	}

	for i := 0; i < len(nums); i++ {
		// https://programmercarl.com/0047.%E5%85%A8%E6%8E%92%E5%88%97II.html#%E6%8B%93%E5%B1%95
		// 如果同一树层(遍历层)nums[i - 1]使用过则直接跳过,树枝也就是递归层可以重复选取
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] == false {
			continue
		}
		if !visited[i] {
			path = append(path, nums[i])
			visited[i] = true
			backtrace(nums, path, visited)
			path = path[:len(path)-1]
			delete(visited, i)
		}
	}
}
