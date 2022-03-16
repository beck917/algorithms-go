package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	mid := slow
	right := reverse(mid)

	cur := head
	for cur != mid {
		if cur.Val != slow.Val {
			return false
		}
		cur = cur.Next
		slow = slow.Next
	}

	if slow != nil {
		return false
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	var next *ListNode

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
