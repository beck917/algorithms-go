package main

import "fmt"

func main() {
	matrix := make([][]int, 3)
	matrix[0] = []int{1, 2, 3}
	matrix[1] = []int{4, 5, 6}
	matrix[2] = []int{7, 8, 9}
	rotate(matrix)

	for _, v := range matrix {
		fmt.Println(v)
	}
}
func rotate(matrix [][]int) {
	for _, nums := range matrix {
		revert(nums)
	}

	for i := 0; i < len(matrix); i++ {
		k := i
		for j := len(matrix[i]) - i - 2; j >= 0; j-- {
			//fmt.Println(matrix[i][j], matrix[k+1][len(matrix[i])-i-1])
			t := matrix[i][j]
			matrix[i][j] = matrix[k+1][len(matrix[i])-i-1]
			matrix[k+1][len(matrix[i])-i-1] = t
			//fmt.Println(matrix[i][j], matrix[k+1][len(matrix[i])-i-1])
			k++
		}
	}

}

func revert(nums []int) {
	left := 0
	right := len(nums) - 1

	for left <= right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
