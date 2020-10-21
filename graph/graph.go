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

var paths []int
var passed map[int]bool

func (graph *Graph) DFS(s int, t int) {
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
	}
}
