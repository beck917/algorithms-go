package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 7))
}

/**
重点是建立一个map,存储遍历过的数据
然后用target(7)-v(当前数据),看map里是否存在,存在就返回数据
*/
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
