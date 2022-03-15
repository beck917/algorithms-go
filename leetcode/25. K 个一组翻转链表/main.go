package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/reverse-nodes-in-k-group/discuss/183356/Java-O(n)-solution-with-super-detailed-explanation-and-illustration
// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7
// step3: 0 (pointer) -> 1 <- 2 <- 3 (prev)    4 (curr) -> 5 -> 6 -> 7
// 0 -> 3 -> 2 -> 1 (pointer) -> 4 -> 5 -> 6 -> 7
/**
看连接文章就可以了,这里简单说一下

这个题目是个困难的是题目,但是只要明白了方法,编写出正确的代码并不困难
首先要判断是否满足k个节点,不满足的话就直接调出loop

满足的话,就按照反转链表的方式,反转前面k个
0 (pointer) -> 1 <- 2 <- 3 (prev) ,注意,这里反转不反转point

tail := pointer.Next
然后将prev接到point上,pointer.next = prev
tail.next = cur 将point next(1)的next指向4
point等于tail(1)
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{0, head}
	pointer := dummyHead
LOOP:
	for pointer != nil {
		// 先判断剩余是否足够k个节点
		node := pointer.Next
		for i := 0; i < k; i++ {
			if node == nil {
				break LOOP
			}
			node = node.Next
		}

		cur := pointer.Next
		var next *ListNode
		var prev *ListNode
		for i := 0; i < k; i++ {
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		// 基本思路是将指针链接到3，将1链接到4: 这里需要注意，需要将pointer链接到prev（3），这样最后才能返回正确的结果
		tail := pointer.Next
		tail.Next = cur
		pointer.Next = prev
		pointer = tail
	}

	return dummyHead.Next
}
