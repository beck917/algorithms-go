package main

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
