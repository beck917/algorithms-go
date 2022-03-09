package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	slow := head

	i := 0
	for fast != nil {
		fast = fast.Next
		if i >= k {
			slow = slow.Next
		}
		i++
	}

	return slow
}
