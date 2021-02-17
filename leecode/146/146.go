package main

import "fmt"

type LRUCache struct {
	cache    map[int]*Node
	capacity int
	list     *List
}

type List struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	next  *Node
	prev  *Node
	key   int
	value int
}

func (list *List) moveToFront(node *Node) {
	list.remove(node)
	list.pushFront(node)
}

func (list *List) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	//node = nil
	list.size--
}

func (list *List) pushFront(node *Node) {
	node.next = list.head.next
	node.prev = list.head
	list.head.next.prev = node
	list.head.next = node
	list.size++
}

func (list *List) removeEnd() {
	node := list.tail.prev
	list.remove(node)
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:    make(map[int]*Node),
		capacity: capacity,
		list: &List{
			head: &Node{},
			tail: &Node{},
			size: 0,
		},
	}
	lru.list.head.next = lru.list.tail
	lru.list.tail.prev = lru.list.head

	return lru
}

func (this *LRUCache) Get(key int) int {
	var (
		node *Node
		ok   bool
	)

	if node, ok = this.cache[key]; !ok {
		return -1
	}

	this.list.moveToFront(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := &Node{key: key, value: value}
		this.cache[key] = node

		if this.list.size >= this.capacity {
			delete(this.cache, this.list.tail.prev.key)
			this.list.removeEnd()
		}
		this.list.pushFront(node)
	} else {
		node := this.cache[key]
		node.value = value
		this.list.moveToFront(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lru := Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(3))
}
