package main

import (
	"fmt"
)

func main() {
	nums := []int{-3, -1}
	fmt.Println(sortedSquares(nums))
}

// 类似归并排序做法，需要临时数组存储，先找到基准线
func sortedSquares(nums []int) []int {
	// 获取基准数据
	index := -1
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			index = i
		} else {
			break
		}
	}

	tmpNums := make([]int, 0)
	j := index + 1
	for i := index; i >= 0 || j < n; {
		// 将对比后，剩余的较大的数据放入数组
		if i < 0 {
			tmpNums = append(tmpNums, nums[j]*nums[j])
			j++
			continue
		}

		if j >= n {
			tmpNums = append(tmpNums, nums[i]*nums[i])
			i--
			continue
		}

		if nums[i]*nums[i] <= nums[j]*nums[j] {
			tmpNums = append(tmpNums, nums[i]*nums[i])
			i--
		} else {
			tmpNums = append(tmpNums, nums[j]*nums[j])
			j++
		}
	}

	return tmpNums
}
