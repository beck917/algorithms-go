package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{0, head}
	//var tailNode *ListNode

	i := 1
	j := 1
	cur := dummyNode
	var startNode *ListNode
	var endNode *ListNode
	for cur != nil {
		if i == left {
			startNode = cur
		}
		if j == right {
			/**
			if cur.Next == nil {
				tailNode = &ListNode{0, nil}
				cur.Next = tailNode
			}
			*/
			endNode = cur.Next
			cur.Next = nil
		}

		cur = cur.Next
		i++
		j++
	}

	cur = startNode.Next
	var next *ListNode
	var prev *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	startNode.Next.Next = endNode
	startNode.Next = prev

	return dummyNode.Next
}
