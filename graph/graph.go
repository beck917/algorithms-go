package main

import "container/list"

// Graph 图的数据结构
type Graph struct {
	v   int // No. of vertices
	adj []*list.List
}

func main() {

}

// New 初始化graph
func (graph *Graph) New(v int) {
	graph.v = v
	for i := 0; i < v; i++ {
		graph.adj[i] = list.New()
	}
}

func (graph *Graph) addEdge(s int, t int) {
	graph.adj[s].PushBack(t)
	graph.adj[t].PushBack(s)
}

var nodePath *Node
var visited map[int]bool
var found bool

// Node 用链表存储遍历过的节点, 这样在回溯的时候,可以将next覆盖掉
type Node struct {
	next  *Node
	value int
}

// DFS test
func (graph *Graph) DFS(s int, t int) {
	nodePath := &Node{
		value: s,
	}
	graph.recurDfs(s, t)
}

func (graph *Graph) recurDfs(s int, t int) {
	visited[node.Value.(int)] = true

	for node := graph.adj[s].Front(); node != nil; node = node.Next() {
		if passed[node.Value.(int)] != true {
			return
		}
		if node.Value != t {
			if passed[node.Value.(int)] != true {
				paths = append(paths, node.Value.(int))
				passed[node.Value.(int)] = true
			}
		} else {
			paths = append(paths, node.Value.(int))
			return
		}
		if !visited[node.Value.(int)] {

			recurDfs(node.Value.(int), t)
		}

	}

}
