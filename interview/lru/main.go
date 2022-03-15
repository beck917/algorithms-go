package main

import "fmt"

func main() {
	lru := NewLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(1))
}

type LRUCache struct {
	cache    map[int]*Node
	list     *ListNode
	capacity int
}

type ListNode struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func (list *ListNode) moveToFront(node *Node) {
	list.remove(node)
	list.pushFront(node)
}

func (list *ListNode) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	list.size--
}

func (list *ListNode) pushFront(node *Node) {
	node.next = list.head.next
	node.prev = list.head

	list.head.next.prev = node
	list.head.next = node
	list.size++
}

func (list *ListNode) removeEnd() {
	list.remove(list.tail.prev)
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache: make(map[int]*Node),
		list: &ListNode{
			head: &Node{},
			tail: &Node{},
		},
		capacity: capacity,
	}

	lru.list.head.next = lru.list.tail
	lru.list.head.prev = lru.list.head

	return lru
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.list.moveToFront(node)
		return node.value
	}

	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		lru.list.moveToFront(node)
		node.value = value

		return
	}

	if lru.capacity < lru.list.size {
		delete(lru.cache, lru.list.tail.prev.key)
		lru.list.remove(lru.list.tail.prev)
	}

	node := &Node{
		key:   key,
		value: value,
	}
	lru.list.pushFront(node)
	lru.cache[key] = node
}
