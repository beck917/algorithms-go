package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Printf("%p 1 \n", nums)
	nums = append(nums, 1)
	fmt.Printf("%p 4 \n", nums)
	results := [][]int{nums}
	nums = nums[:len(nums)-1]
	change(nums)
	fmt.Println(nums)
	fmt.Printf("%p 3 \n", nums)
	fmt.Println(results)

	a := make([]int, 0, 100)
	a = append(a, 1, 2, 3)
	foo(a)
	//fmt.Println(a)
}

func foo(nums []int) {
	nums = append(nums, 2)
}

func change(nums []int) {
	nums = append(nums, 2)
	fmt.Printf("%p 2 \n", nums)
	//fmt.Println(nums)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}

	point := dummyHead

	for point != nil {
		if point.Next == nil || point.Next.Next != nil {
			break
		}

		var prev *ListNode
		var next *ListNode
		cur := point.Next
		for i := 0; i < 2; i++ {
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		next = point.Next
		point.Next = prev
		next.Next = cur
		point = next
	}

	return dummyHead.Next
}
