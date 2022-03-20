package main

func majorityElement(nums []int) int {
	visited := make(map[int]int, 0)

	for i := 0; i < len(nums)-1; i++ {
		visited[nums[i]]++
	}

	prev := 0
	result := 0
	for k, v := range visited {
		if v > prev {
			result = k
		}
	}

	return result
}
