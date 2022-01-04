package main

import (
	//"container/list"
	"fmt"
)

func main() {
	fmt.Println("main")

	//list.New()

	//list := []int{5, 2, 1, 4, 3, 6, 8, 7}
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(binarySearch(list, 8))
}

func binarySearch(list []int, target int) int {
	left := 0
	right := len(list) - 1
	mid := right / 2
	for left <= right {
		if target < list[mid] {
			right = mid - 1
		} else if target > list[mid] {
			left = mid + 1
		} else {
			return mid
		}
		mid = (left + right) / 2
	}
	return -1
}

func quickSort() {

}
