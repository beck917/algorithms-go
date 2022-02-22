package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/reverse-nodes-in-k-group/discuss/183356/Java-O(n)-solution-with-super-detailed-explanation-and-illustration
// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7
// step3: 0 (pointer) -> 1 <- 2 <- 3 (prev)    4 (curr) -> 5 -> 6 -> 7
// 0 -> 3 -> 2 -> 1 (pointer) -> 4 -> 5 -> 6 -> 7
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
