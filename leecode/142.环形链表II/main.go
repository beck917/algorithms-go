package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {

		}
	}
}
