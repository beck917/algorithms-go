package main

import "fmt"

func main() {
	nums := []int{5, 4, 6, 2}
	fmt.Println(permute(nums))
}

var results [][]int

/**
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func permute(nums []int) [][]int {
	n := len(nums)
	results = make([][]int, 0)
	used := make([]bool, n)
	for i, _ := range nums {
		used[i] = false
	}
	path := make([]int, 0)
	backTrack(nums, used, path)

	return results
}

func backTrack(nums []int, used []bool, path []int) {
	if len(path) == len(nums) {
		// 因为我们要输出最终的结果集,path是个数组,在回溯过程中,会被修改,这里我们需要将path拷贝一份
		p := make([]int, len(path))
		copy(p, path)
		results = append(results, p)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == false {
			//fmt.Println(used, path, nums[i])
			used[i] = true

			path = append(path, nums[i])

			backTrack(nums, used, path)
			// 为什么输出[1,2,3]后, 回溯到[1]之后, 这里的循环还在继续, 运行了[1,2]之后, 这时轮到3了
			// 这时候就需要将2踢出, used[1]也标记为false
			path = path[0 : len(path)-1]
			used[i] = false
		}
	}
}
