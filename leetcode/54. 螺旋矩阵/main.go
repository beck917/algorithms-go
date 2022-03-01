package main

import "fmt"

func main() {
	matrix := make([][]int, 0)
	matrix = append(matrix, []int{1, 2, 3})
	matrix = append(matrix, []int{4, 5, 6})
	matrix = append(matrix, []int{7, 8, 9})
	matrix = append(matrix, []int{10, 11, 12})

	fmt.Println(spiralOrder(matrix))
}

// [[1,2,3],[4,5,6],[7,8,9],[10,11,12]]
// [1,2,3,6,9,12,11,10,7,4,5,8]
func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	for len(matrix) != 0 {
		tmpMatrix := make([][]int, 0)
		tmpResult := make([]int, 0)
		for k, v := range matrix {
			if k == 0 {
				for i := 0; i < len(v); i++ {
					result = append(result, v[i])
				}
			} else if k == len(matrix)-1 {
				for i := len(v) - 1; i >= 0; i-- {
					result = append(result, v[i])
				}
			} else {
				tmp := make([]int, 0)
				for i := 1; i < len(v)-1; i++ {
					tmp = append(tmp, v[i])
				}
				tmpMatrix = append(tmpMatrix, tmp)
				if len(v) >= 1 {
					result = append(result, v[len(v)-1])
				}
				if len(v) > 1 {
					tmpResult = append(tmpResult, v[0])
				}
			}
		}

		for i := len(tmpResult) - 1; i >= 0; i-- {
			result = append(result, tmpResult[i])
		}

		matrix = tmpMatrix
	}

	return result
}
