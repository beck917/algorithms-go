package main

/**
输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
var results [][]int

func combine(n int, k int) [][]int {
	results = make([][]int, 0)
	visited := make(map[int]bool, n)
	nums := make([]int, 0)
	/**
	for i := 1; i <= n; i++ {
		nums = []int{i}
		visited[i] = true
		backtrace(i, nums, n, k, visited)
	}*/
	backtrace(1, nums, n, k, visited)

	return results
}

func backtrace(index int, nums []int, n, k int, visited map[int]bool) {
	if len(nums) == k {
		p := make([]int, k)
		copy(p, nums)
		results = append(results, p)
		return
	}

	// 剪枝, 当前的数量+剩余的数量(-1是因为前面+1了)
	if len(nums)+(n-(index-1)) < k {
		return
	}

	for i := index; i <= n; i++ {
		if len(nums)+1 > k {
			return
		}
		if !visited[i] {
			nums = append(nums, i)
			//visited[i] = true
			backtrace(i+1, nums, n, k, visited) // 这里传递了+1,不需要visited了,因为前面的数不需要了
			nums = nums[:len(nums)-1]
			//delete(visited, i)
		}
	}
}
