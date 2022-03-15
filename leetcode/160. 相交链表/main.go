package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
这个题我的方法比较取巧

先找到两个链表长度一致的地方,然后从这个地方同时遍历链表,如果相等就相交
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmpHeadA := headA
	tmpHeadB := headB

	ai := 0
	for tmpHeadA != nil {
		tmpHeadA = tmpHeadA.Next
		ai++
	}
	bi := 0
	for tmpHeadB != nil {
		tmpHeadB = tmpHeadB.Next
		bi++
	}

	if ai > bi {
		for i := 0; i < ai-bi; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < ai-bi; i++ {
			headB = headB.Next
		}
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
