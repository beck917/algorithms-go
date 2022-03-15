package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{0, nil}
	dummyHead := &ListNode{0, l3}

	var carry int
	var sum int
	cur := dummyHead
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		val := n1 + n2 + carry
		sum = val % 10
		carry = val / 10
		cur.Next = &ListNode{sum, nil}
		cur = cur.Next

	}

	if carry != 0 {
		cur.Next = &ListNode{carry, nil}
	}

	return dummyHead.Next
}
