package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 很简单的题目,做法和层序遍历一样,只需要在i/2==0的时候做特殊判断就可以了
// 此题目为中等难度, 思路虽然简单,但很容易出错
func zigzagLevelOrder(root *TreeNode) [][]int {
	levelNodes := make([][]int, 0)

	if root == nil {
		return levelNodes
	}

	levelNodes = append(levelNodes, []int{root.Val})

	parentsNodes := []*TreeNode{root}
	index := 0
	for len(parentsNodes) > 0 {
		// 存储遍历过的父节点,这里只需要同样的顺序存储就可以了,因为在下面遍历节点的时候回改变遍历顺序
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
