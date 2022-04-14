package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	}

	j := len(matrix[0]) - 1
	i := len(matrix) - 1
	for i >= 0 && j >= 0 {
		if target < matrix[i][j] {
			j--
		} else if target > matrix[i][j] {
			i--
		} else {
			return true
		}
	}

	return false
}
