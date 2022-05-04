package main

func findDiagonalOrder(mat [][]int) []int {
	i := 0
	j := 0

	var nums []int
	var flag bool
	for {
		if i >= 0 && i <= len(mat)-1 && j >= 0 && j <= len(mat[0])-1 {
			nums = append(nums, mat[i][j])

			if !flag {
				i--
				j++
			} else {
				i++
				j--
			}
		} else {
			if i < 0 {
				i = 0
			}
			if j < 0 {
				j = 0
			}
			if i > len(mat)-1 {
				i = len(mat) - 1
			}
			if j > len(mat[0])-1 {
				j = len(mat[0]) - 1
			}
		}

		if i == len(mat)-1 && j == len(mat[0])-1 {
			break
		}
	}

	return nums
}
