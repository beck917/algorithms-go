package main

import "fmt"

func main() {
	nums1 := []int{2, 3, 0, 0}
	nums2 := []int{2, 4}

	merge(nums1, 2, nums2, 2)
	fmt.Println(nums1)
}

// 从后往前遍历
func merge(nums1 []int, m int, nums2 []int, n int) {
	len1 := len(nums1)
	len2 := len(nums2)
	i := len1 - len2 - 1
	j := len2 - 1

	for k := len1 - 1; k >= 0; k-- {
		if j < 0 {
			// add left data
			nums1[k] = nums1[i]
			i--
			continue
		}

		if i < 0 {
			nums1[k] = nums2[j]
			j--
			continue
		}

		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
