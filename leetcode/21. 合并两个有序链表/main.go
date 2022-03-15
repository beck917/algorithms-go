package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
建立一个新的链表

注意将剩余的那个链表,接到newlist后面就行了,不需要遍历了
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newListTmp := &ListNode{0, nil}
	newList := newListTmp
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			newList.Next = list1
			list1 = list1.Next
		} else {
			newList.Next = list2
			list2 = list2.Next
		}
		newList = newList.Next
	}

	if list1 == nil {
		for list2 != nil {
			newList.Next = list2
			list2 = list2.Next
			newList = newList.Next
		}
	}

	if list2 == nil {
		for list1 != nil {
			newList.Next = list1
			list1 = list1.Next
			newList = newList.Next
		}
	}

	return newListTmp.Next
}
