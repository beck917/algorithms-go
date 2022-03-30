package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
卧槽.这个题目用归并排序,绝对是困难级别

首先归并排序即便是数组都很难
这里我们还是理一下思路,帮助记忆和背诵
第一步,确定sort递归模板
func sort(head, tail *ListNode) *ListNode {
传入头尾节点,不断遍历
	mid := slow
	list1 := sort(head, mid)
	list2 := sort(mid, tail)
    快慢指针寻找到中间节点, 这里是一个分拆的过程,将大链表分拆成小链表

    return mergeSort(list1, list2)
    最后将小链表,合并,这里就是合并两个排序链表
*/
func mergeSort(list1, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}
	tmplist, cur1, cur2 := dummyHead, list1, list2
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			tmplist.Next = cur1 //注意这里不是新建链表
			cur1 = cur1.Next
		} else {
			tmplist.Next = cur2
			cur2 = cur2.Next
		}
		//tmplist.Next = &ListNode{val, nil}
		tmplist = tmplist.Next
	}

	if cur1 != nil {
		tmplist.Next = cur1
	}

	if cur2 != nil {
		tmplist.Next = cur2
	}

	return dummyHead.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		// 左边标记为nil
		head.Next = nil
		return head
	}

	fast, slow := head, head
	// 快慢指针寻找中间节点
	for fast != tail {
		fast = fast.Next
		slow = slow.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	list1 := sort(head, mid)
	list2 := sort(mid, tail)

	return mergeSort(list1, list2)
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}
