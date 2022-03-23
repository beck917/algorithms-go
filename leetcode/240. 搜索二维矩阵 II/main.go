package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	}

	left := 0
	up := 0
	right := len(matrix[0])
	down := len(matrix)

	for left <= right || up <= down {
		midh := (left + right) >> 1
		midv := (up + down) >> 1

		if matrix[]
	}
}
