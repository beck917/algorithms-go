package main

import "fmt"

// 所谓的快慢指针法
// 其实就是遍历的时候，两个游标累加速度不同，和快排，归并排序的算法类似
func removeElement2(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

func main() {
	nums := []int{-1, 0, 5, 5, 9, 12}
	len := removeElement(nums, 5)

	fmt.Println(len)
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	j := 0
	count := len(nums)
	for i := 0; i < count; i++ {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	return j
}
