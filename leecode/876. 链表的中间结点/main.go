package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil {
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
		slow = slow.Next
	}

	return slow
}
