package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{math.MaxInt64, head}

	cur := dummyHead
	var prev *ListNode
	for cur != nil {
		if prev != nil && cur.Val == prev.Val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}

	return dummyHead.Next
}
