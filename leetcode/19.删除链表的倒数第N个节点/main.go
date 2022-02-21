package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针方式
// 快指针先开始走，慢指针等快指针走了n个节点后才开始走
// len-n = slow, 这样快指针走完，慢指针刚好到要删除的节点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{0, head}

	fast := dummyHead
	slow := dummyHead

	i := 0
	for fast != nil {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}
		i++
	}

	slow.Next = slow.Next.Next

	return dummyHead.Next
}
