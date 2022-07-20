package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{math.MaxInt64, head} // 这里要注意用最大值

	cur := dummyHead
	var prev *ListNode
	for cur != nil {
		if prev != nil && cur.Val == prev.Val {
			prev.Next = cur.Next
		} else {
			prev = cur //这里要注意,prev只有在链表没改变的时候,改变值,因为要保留prev连接到下一个
		}
		cur = cur.Next
	}

	return dummyHead.Next
}
