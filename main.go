package main

import (
	"sync"
	"sync/atomic"
	//"container/list"
	"fmt"
)

func main() {
	fmt.Println("main")

	arr := []int{6, 3, 1, 7, 4, 9, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	index := left
	for j := left; j < right; j++ {
		// 关键点，在于将小的数据，移到左边去，而不是将大的数据移到右边
		// 其实也是一种快慢指针，index从0开始，如果0索引位置<pivot,那么就是和自己交换
		if arr[j] < pivot {
			arr[index], arr[j] = arr[j], arr[index]
			index++
		}
	}
	// 交换pivot和 index
	arr[index], arr[right] = arr[right], arr[index]

	return index
}

// quicksort() =
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	index := partition(arr, left, right)
	quickSort(arr, left, index-1)
	quickSort(arr, index+1, right)
}

type ListNode struct {
	val  int
	next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.next
		// 这里的关键点是cur next指向prev， 而在开始时prev正好为nil
		cur.next = prev
		prev = cur
		cur = cur.next
	}

	return prev
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

type Graph struct {
	v int // 定点的数量
	adj [][]int
}

func (graph *Graph) New(v int) {
	graph.v = v
	for i := 0; i < v; i++ {
		graph.adj[i] = make([]int, 0))
	}
}

func (graph *Graph) addEdge(s, t int) {
	graph.adj[s] = append(graph.adj[s], t)
	graph.adj[t] = append(graph.adj[t], s)
}

type Once struct {
	mu sync.Mutex
	flag atomic.Int32
}

func (once *Once) Do(fn func()) {
	flag := atomic.Get(&once.flag)
	
	if flag == 0 {
		once.mu.Lock()
		defer once.mu.Unlock()
		flag = atomic.Get(&once.flag)
		if flag != 0 {
			defer atomic.Set(&once.flag, 1)
			fn()
		}
	}
}
