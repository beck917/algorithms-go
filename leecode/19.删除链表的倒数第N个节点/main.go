package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{0, head}

	fast := dummyHead
	slow := dummyHead

	i := 0
	for fast != nil {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}
		i++
	}

	slow.Next = slow.Next.Next

	return dummyHead.Next
}
