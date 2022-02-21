package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		// 需要公式推到，非人类思维可以直接理解。。。
		if fast == slow {
			for slow.Next
		}
	}

	return false
}
