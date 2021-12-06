package main

import (
	"container/list"
	"fmt"
)

// Graph 图的数据结构
type Graph struct {
	v   int // No. of vertices
	adj []*list.List
}

func main() {
	g := &Graph{}
	g.New(6)
	g.addEdge(0, 1)
	g.addEdge(1, 2)
	g.addEdge(2, 3)
	g.addEdge(3, 5)
	g.addEdge(2, 4)
	g.DFS(0, 4)
	//g.BFS(0, 4)
}

// New 初始化graph
func (graph *Graph) New(v int) {
	graph.v = v
	graph.adj = make([]*list.List, 6)
	for i := 0; i < v; i++ {
		graph.adj[i] = list.New()
	}
}

func (graph *Graph) addEdge(s int, t int) {
	graph.adj[s].PushBack(t)
	graph.adj[t].PushBack(s)
}

var nodePath *Node
var visited map[int]bool = make(map[int]bool)
var found bool

// Node 用链表存储遍历过的节点, 这样在回溯的时候,可以将next覆盖掉
type Node struct {
	next  *Node
	value int
}

func (l *Node) traversal() {
	for l != nil {
		fmt.Println(l.value)
		l = l.next
	}
}

// DFS test
func (graph *Graph) DFS(s int, t int) {
	nodePath := &Node{
		value: s,
	}
	graph.recurDfs(s, t, nodePath)
	nodePath.traversal()
}

func (graph *Graph) recurDfs(s int, t int, nodePath *Node) {
	fmt.Println("visited node ", s)
	visited[s] = true
	if s == t {
		found = true
		return
	}

	for node := graph.adj[s].Front(); node != nil; node = node.Next() {
		if found == true {
			return
		}

		if !visited[node.Value.(int)] {
			nodePath.next = &Node{value: node.Value.(int)}
			graph.recurDfs(node.Value.(int), t, nodePath.next)
		}

	}
}

// prev 数组，存储遍历过的节点，通过记录前置节点的方式，实际上就是用数组表示链表
// queue 存储即将遍历的子节点
// visted 存储遍历过的节点
func (graph *Graph) BFS(s, t int) {
	if s == t {
		return
	}
	var prev []int = make([]int, graph.v)
	var queue = list.New()
	var visited = make(map[int]bool)

	for i := 0; i < graph.v; i++ {
		prev[i] = -1
	}

	queue.PushBack(s)

	for queue.Len() != 0 {
		elem := queue.Back().Value.(int)

		for node := graph.adj[elem].Front(); node != nil; node = node.Next() {
			cur := node.Value.(int)
			if !visited[cur] {
				prev[cur] = elem
			}

			if cur == t {
				graph.print(prev, s, t)
				return
			}
			visited[cur] = true
			queue.PushBack(cur)
		}
	}
}

func (graph *Graph) print(prev []int, s, t int) {
	if s != t && prev[t] != -1 {
		graph.print(prev, s, prev[t])
	}
	fmt.Println(t)
}
