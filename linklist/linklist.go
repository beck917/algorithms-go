package main

import "fmt"

/**

//ElemType is list type
type ElemType interface{}

//Node is list node
type Node struct {
	Data ElemType
	Next *Node
}

//LinkedList is single list
type LinkedList struct {
	Head *Node
}

*/

type linkList struct {
	next  *linkList
	value int
}

func main() {
	l := &linkList{}
	l.value = 0
	next := l.insert(&linkList{value: 1})
	next = next.insert(&linkList{value: 2})
	next = next.insert(&linkList{value: 3})
	next = next.insert(&linkList{value: 4})
	l = l.reverse()
	l.traversal()
}

func (l *linkList) insert(next *linkList) *linkList {
	l.next = next
	return next
}

func (l *linkList) traversal() {
	for l != nil {
		fmt.Println(l.value)
		l = l.next
	}
}

// 4->3, 3->2, 2->1, 1->0， 0->nil
// 转换next指针，相当于将next变成prev
func (l *linkList) reverse() *linkList {
	cur := l
	var prev *linkList
	var next *linkList
	for cur.next != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	cur.next = prev
	return cur
}
