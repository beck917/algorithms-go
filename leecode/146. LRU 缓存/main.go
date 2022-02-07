package main

type LRUCache struct {
	capacity int
	cache    map[int]*Node
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
	list.size--
}

func (list *List) pushFront(node *Node) {
	node.next = list.head.next
	node.prev = list.head
	list.head.next.prev = node
	list.head.next = node
	list.size++
}

func (list *List) removeTail() {
	list.remove(list.tail.prev)
}

func Constructor(capacity int) LRUCache {
	list := &List{
		head: &Node{},
		tail: &Node{},
		size: 0,
	}

	list.head.next = list.tail
	list.tail.prev = list.head
	return LRUCache{
		list:     list,
		capacity: capacity,
		cache:    make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.list.moveToFront(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		this.list.moveToFront(node)
		node.value = value
	} else {
		if this.list.size >= this.capacity {
			delete(this.cache, this.list.tail.prev.key)
			this.list.removeTail()
		}

		node := &Node{nil, nil, key, value}
		this.list.pushFront(node)
		this.cache[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
