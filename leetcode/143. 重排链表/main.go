package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]
*/
func reorderList(head *ListNode) {
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
