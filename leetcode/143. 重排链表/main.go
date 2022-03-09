package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	fast := head
	slow := head
	i := 0
	// 获取链表中间节点
	for fast != nil {
		fast = fast.Next
		if i%2 == 0 {
			slow = slow.Next
		}
		i++
	}

	mid := slow
	left := head
	right := mid.Next

	// 反转右边链表
	right = reverse(right)

	// 合并链表
	cur := left
	for cur != mid {
		next := cur.Next
		rightnext := right.Next
		cur.Next = right
		right.Next = next

		cur = next
		right = rightnext
	}
	mid.Next = nil
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var next, prev *ListNode

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

/**
输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]
*/
func reorderList2(head *ListNode) {
	if head.Next == nil || head.Next.Next == nil {
		return
	}

	tmpHead := head
	for tmpHead.Next == nil || tmpHead.Next.Next != nil {
		tmpHead = changeHead(tmpHead)
	}
}

func changeHead(head *ListNode) *ListNode {
	tmpHead := head
	cur := head
	next := head.Next
	var prev *ListNode
	for {
		if cur.Next == nil {
			tmpHead.Next = cur // 1->5
			//tmpHead = cur
			cur.Next = next // 5->2
			tmpHead = next  // 2

			prev.Next = nil
			break
		}
		prev = cur
		cur = cur.Next
	}

	return tmpHead
}
