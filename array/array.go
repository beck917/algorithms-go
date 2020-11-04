package main

import "fmt"

func main() {
	list := []int{5, 2, 1, 4, 3, 6, 8, 7}
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
	fmt.Println(binarySearch(list, 5))
}

func binarySearch(list []int, target int) int {
	left := 0
	right := len(list) - 1

	for right >= left {
		mid := (left + right) >> 1
		if target > list[mid] {
			left = mid + 1
		} else if target < list[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

// quickSort1
func quickSort(list []int, start, end int) {
	if start >= end {
		return
	}

	pivot := end

	index := start - 1

	for i := start; i < end; i++ {
		// 如果i数据小于对比数据pivot，则i和index交换，将小的数据移到左边
		if list[i] < list[pivot] {
			index++
			// 交换数据
			list[i], list[index] = list[index], list[i]
		}
	}

	// 将对比数据pivot移到index+1的位置
	mid := index + 1
	list[pivot], list[mid] = list[mid], list[pivot]

	// mid的数组位置已经固定，不需要参与继续的递归排序了
	quickSort(list, start, mid-1)
	quickSort(list, mid+1, end)
}
