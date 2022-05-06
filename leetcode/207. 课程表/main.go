package main

type Node struct {
	indeg int     // 所有指向这个节点的课程节点之和(被指向)
	adj   []*Node // 出度 邻节点(指向)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	nodeMap := make(map[int]*Node, 0)

	// 初始化所有
	for i := 0; i < numCourses; i++ {
		nodeMap[i] = &Node{
			0,
			make([]*Node, 0),
		}
	}

	for _, v := range prerequisites {
		nodeMap[v[1]].adj = append(nodeMap[v[1]].adj, nodeMap[v[0]]) // 指向节点, 学这个课程前需要学习哪些课程
		nodeMap[v[0]].indeg++                                        // 被指向,需要学习的前置课程数量累加
	}

	// 类似bfs, 维护parentNodes, 讲没有前置课程的节点放入
	var parentNodes []*Node
	for i := 0; i < numCourses; i++ {
		if nodeMap[i].indeg == 0 {
			parentNodes = append(parentNodes, nodeMap[i])
		}
	}

	count := 0 //最后判断遍历过的节点数和numCourses是否一致
	// 因为如果有环, 临节点必然不能被加入队列
	for len(parentNodes) != 0 {
		var tmpNodes []*Node
		for _, node := range parentNodes {
			count++
			// 遍历邻接点, 将前置课程节点移除,看是否还需要前置
			for _, adj := range node.adj {
				adj.indeg--
				if adj.indeg == 0 {
					tmpNodes = append(tmpNodes, adj)
				}
			}
		}
		parentNodes = tmpNodes
	}
	return count == numCourses
}
