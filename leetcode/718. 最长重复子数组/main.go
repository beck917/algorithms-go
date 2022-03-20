package main

import "fmt"

func main() {
	fmt.Println(findLength([]int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}))
}

/**
输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
输出：3
解释：长度最长的公共子数组是 [3,2,1] 。
*/
func findLength(nums1 []int, nums2 []int) int {
	max := 0
	k := 0
	for i := 0; i < len(nums1); i++ {
		k = i
		tmp := 0
		for j := 0; k < len(nums1) && j < len(nums2); j++ {
			fmt.Println(i, k, j, nums1[k], nums2[j])
			if nums1[k] == nums2[j] {
				k++
				tmp++
			} else {
				if tmp > max {
					max = tmp
				}
				tmp = 0
			}
		}

		if tmp > max {
			max = tmp
		}
	}
	return max
}
