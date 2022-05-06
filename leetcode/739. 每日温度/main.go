package main

// 单调栈,还是比较简单的, 注意栈里按照递增存储索引即可
// https://leetcode-cn.com/problems/daily-temperatures/solution/leetcode-tu-jie-739mei-ri-wen-du-by-misterbooo/
func dailyTemperatures(temperatures []int) []int {
	var stack []int
	results := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			results[prevIndex] = i - prevIndex
		}

		stack = append(stack, i)
	}

	return results
}
