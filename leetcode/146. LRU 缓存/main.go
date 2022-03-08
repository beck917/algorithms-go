package main

/**
首先理解lru的目的,存储特定容量的数据,将最近没有访问的数据淘汰掉.这样的话就需要一个
list双向链表来淘汰数据
lruCache 需要capacity 容量, cache的map用来直接找到对应的node,
list是一个双向链表,用来淘汰数据,后head,tail和size
list结构,为了方便数据的删减, 使用head和tail存储头尾节点
node 存储prev,next,key,value

我们看一下需要的方法,基本的需要get,put
get方法从cache获取数据,成功的化,需要将此节点的链表数据放到最前
put方法从cache获取数据,如果有数据,则更新数据,并将数据放到最前,没有数据的话,
则需要判断链表容量,delete()方法删除cache的数据,并用removetail删除尾节点
然后将数据放在链表最前pushFront,并更新缓存值

下面说一下get和put依赖的几个核心方法
movetofront, 这个方法就是remove()删除当前节点,然后调用pushfront
remove, 这个删除很简单,就是嫁接下链表,记得要size--
pushFront, 这个就是把当前节点的next和prev设置下,然后设置下头结点的下一个节点的的next和prev
removetail,这个只需要删除tail的prev就行了,调用remove

这里还要说一下类的初始化防范,因为要初始化list的tail和head节点
首先将list的head和tail建立&Node{}, 这里需要注意的是,需要将head和tail收尾相连
	list.head.next = list.tail
	list.tail.prev = list.head
*/
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
