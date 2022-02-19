package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	levelNodes := make([][]int, 0)

	if root == nil {
		return levelNodes
	}

	levelNodes = append(levelNodes, []int{root.Val})

	parentsNodes := []*TreeNode{root}
	index := 0
	for len(parentsNodes) > 0 {
		tmpNodes := []*TreeNode{}
		for i := 0; i < len(parentsNodes); i++ {
			pnode := parentsNodes[i]
			if pnode.Left != nil {
				tmpNodes = append(tmpNodes, pnode.Left)
			}
			if pnode.Right != nil {
				tmpNodes = append(tmpNodes, pnode.Right)
			}
		}
		nodes := []int{}
		if index%2 == 0 {
			for i := len(parentsNodes) - 1; i >= 0; i-- {
				pnode := parentsNodes[i]
				if pnode.Right != nil {
					nodes = append(nodes, pnode.Right.Val)
				}
				if pnode.Left != nil {
					nodes = append(nodes, pnode.Left.Val)
				}
			}
			if len(nodes) != 0 {
				levelNodes = append(levelNodes, nodes)
			}
		} else {
			for i := 0; i < len(parentsNodes); i++ {
				pnode := parentsNodes[i]
				if pnode.Left != nil {
					nodes = append(nodes, pnode.Left.Val)
				}
				if pnode.Right != nil {
					nodes = append(nodes, pnode.Right.Val)
				}
			}
			if len(nodes) != 0 {
				levelNodes = append(levelNodes, nodes)
			}
		}
		index++
		parentsNodes = tmpNodes
	}

	return levelNodes
}
