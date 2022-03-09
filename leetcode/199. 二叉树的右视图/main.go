package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var path []int
var pathMap map[int]int

func rightSideView(root *TreeNode) []int {
	path = []int{}
	pathMap = make(map[int]int)
	dfs(root, 0)

	for i := 0; i < len(pathMap); i++ {
		path = append(path, 0)
	}

	for k, v := range pathMap {
		path[k] = v
	}

	return path
}

func dfs(node *TreeNode, level int) {
	if node == nil {
		return
	}

	pathMap[level] = node.Val
	dfs(node.Left, level+1)
	dfs(node.Right, level+1)
}

func bfs(node *TreeNode) {
	parentNodes := []*TreeNode{node}

	for len(parentNodes) != 0 {
		tmpNodes := []*TreeNode{}
		for k, v := range parentNodes {
			if k == len(parentNodes)-1 {
				path = append(path, v.Val)
			}
			if v.Left != nil {
				tmpNodes = append(tmpNodes, v.Left)
			}
			if v.Right != nil {
				tmpNodes = append(tmpNodes, v.Right)
			}
		}
		parentNodes = tmpNodes
	}
}
