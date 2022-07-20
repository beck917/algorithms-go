package main

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	nodelist *NodeList
}

type NodeList struct {
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

func (nl *NodeList) moveToFront(node *Node) {
	nl.remove(node)
	nl.pushToFront(node)
}

func (nl *NodeList) pushToFront(node *Node) {
	node.prev = nl.head
	node.next = nl.head.next
	nl.head.next.prev = node
	nl.head.next = node
	nl.size++
}

func (nl *NodeList) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	nl.size--
}

func (nl *NodeList) removeTail() {
	nl.remove(nl.tail.prev)
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		nodelist: &NodeList{
			&Node{},
			&Node{},
			0,
		},
	}

	lru.nodelist.head.next = lru.nodelist.tail
	lru.nodelist.tail.prev = lru.nodelist.head

	return lru
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.nodelist.moveToFront(v)
		return v.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 判断是否存在key,存在的话更新
	if v, ok := this.cache[key]; ok {
		this.nodelist.moveToFront(v)
		v.value = value
		this.cache[key] = v

		return
	}

	// 如果不存在,则删除最老的
	if this.nodelist.size >= this.capacity {
		delete(this.cache, this.nodelist.tail.prev.key)
		this.nodelist.removeTail()
	}

	node := &Node{
		key:   key,
		value: value,
	}

	this.nodelist.pushToFront(node)

	this.cache[key] = node
}
