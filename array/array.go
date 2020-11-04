package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(binarySearch(list, 2))
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

func quickSort(list []int, left, right int) {
	pivot := list[right]
	index := left - 1

}
