package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}

	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			firstVal := cur.Next.Val

			for cur.Next != nil && cur.Next.Val == firstVal {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
