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
		if fast == slow {
			ptr := head
			for ptr != nil {
				if ptr == slow {
					return ptr
				}
				slow = slow.Next
				ptr = ptr.Next
			}
		}
	}

	return nil
}
