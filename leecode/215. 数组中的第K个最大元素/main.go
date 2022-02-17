package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2

	fmt.Println(findKthLargest2(nums, k))
}

func findKthLargest(nums []int, k int) int {
	return findQuick(nums, 0, len(nums)-1, len(nums)-k)
}

func partition(nums []int, start, end int) int {
	pivot := end
	j := start
	for i := start; i < end; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	nums[pivot], nums[j] = nums[j], nums[pivot]
	return j
}

func findQuick(nums []int, start, end, k int) int {
	//if start >= end {
	//	return 0
	//}

	index := partition(nums, start, end)
	//fmt.Println(index, k, nums[index])
	if index > k {
		return findQuick(nums, start, index-1, k)
	} else if index < k {
		return findQuick(nums, index+1, end, k)
	}

	return nums[index]
}

//-----堆排序方案------------
func findKthLargest2(nums []int, k int) int {
	heapSize := len(nums)
	// buildHeap
	for i := len(nums) / 2; i >= 0; i-- {
		heapify(nums, heapSize, i)
	}

	for i := 0; i < k-1; i++ {
		// 交换堆顶元素和最后一个元素， 这样相当于删除了堆顶
		nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
		heapSize--
		// 因为删除了堆顶，所以要重新构建堆化
		heapify(nums, heapSize, 0)
	}

	return nums[0]

}

// n 数组长度 i 需要堆化的节点
func heapify(nums []int, n, i int) {
	for {
		left, right, maxPos := i*2+1, i*2+2, i
		if left < n && nums[maxPos] < nums[left] {
			maxPos = left
		}
		if right < n && nums[maxPos] < nums[right] {
			maxPos = right
		}
		if i == maxPos {
			break
		}

		nums[i], nums[maxPos] = nums[maxPos], nums[i]
		i = maxPos
	}
}
