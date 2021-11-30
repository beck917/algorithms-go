package main

import "fmt"

func main() {
	list := []int{5, 2, 1, 4, 3, 6, 8, 7}
	//quickSort(list, 0, len(list)-1)
	//fmt.Println(list)
	//fmt.Println(binarySearch(list, 5))
	mergeSort(list, 0, len(list)-1)
	fmt.Println(list)
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

// quickSort
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

// merge(mergeSort(p,q), mergeSort(q+1, r))
// p >= r
func mergeSort(list []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2

	mergeSort(list, start, mid)
	mergeSort(list, mid+1, end)
	merge(list, start, mid, end)
}

func merge(list []int, left, mid, right int) {
	var tempList []int
	i := left
	j := mid + 1
	//k := 0
	for i <= mid && j <= right {
		if list[i] < list[j] {
			// 分片数组里如果值更新，则++继续从这个数组取
			// 这种做法和合并多个有序文件，是一样的
			tempList = append(tempList, list[i])
			i++
		} else {
			tempList = append(tempList, list[j])
			j++
		}
		//k++
	}

	// 遍历剩余的数据，如果有一个数组写入temp完成，那么另外一个数组则没有机会写入、
	if i <= mid {
		for i <= mid {
			tempList = append(tempList, list[i])
			i++
		}
	} else {
		for j <= right {
			tempList = append(tempList, list[j])
			j++
		}
	}

	// 遍历赋值
	l := left
	for i := 0; i < len(tempList); i++ {
		list[l] = tempList[i]
		l++
	}
}
