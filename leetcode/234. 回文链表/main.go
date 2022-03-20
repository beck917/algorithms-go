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

	right := reverse(slow)

	cur := head
	for cur != slow && right != nil {
		if cur.Val != right.Val {
			return false
		}
		cur = cur.Next
		right = right.Next
	}

	if right != nil {
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
