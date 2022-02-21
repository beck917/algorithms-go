package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 7))
}

func twoSum(nums []int, target int) []int {
	targetMap := make(map[int]int, 0)
	for i, v := range nums {
		if i2, ok := targetMap[target-v]; ok {
			return []int{i2, i}
		}
		targetMap[v] = i
	}
	return []int{}
}
