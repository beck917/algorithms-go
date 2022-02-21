package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) traversal() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

// reverseList prev-->新链表, 将最新的节点指向新链表prev, 起到反转的作用
func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func (l *ListNode) reverseList2() *ListNode {
	cur := l
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
