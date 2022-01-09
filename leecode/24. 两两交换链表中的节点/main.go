package main

type ListNode struct {
	val  int
	next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{0, head}
	cur := dummyNode
	//dummyNode.next = cur.next.next

	for cur.next != nil && cur.next.next != nil {
		next1 := cur.next
		next2 := cur.next.next
		// 这个算法的关键点，在于需要将当前节点指向后面交换过的两个节点中的后一个
		// 解决了两个问题，1. dummy节点的next=头节点，2. 后续的cur不会因为后面两个节点的交换而断开
		cur.next = next2
		next1.next = next2.next
		next2.next = next1
		cur = next1
	}

	return dummyNode.next
}
