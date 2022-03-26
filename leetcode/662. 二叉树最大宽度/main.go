package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueNode struct {
	Node  *TreeNode
	Pos   int
	Level int
}

/**
乍一看挺简单,实际上比较难
1. pos要用堆的思想 2. queue要建立新的类型
*/
func widthOfBinaryTree(root *TreeNode) int {
	parentNodes := []*QueueNode{&QueueNode{root, 0, 0}}

	max := 0
	level := 1
	for len(parentNodes) != 0 {
		var tmpNodes []*QueueNode
		tmpMax := 0

		leftPos := 0
		rightPos := 0
		for _, node := range parentNodes {
			if node.Level != level {
				leftPos = node.Pos
				level++
			}

			if node.Node.Left != nil {
				tmpNodes = append(tmpNodes, &QueueNode{node.Node.Left, node.Pos * 2, level})
				rightPos = node.Pos
			}

			if node.Node.Right != nil {
				tmpNodes = append(tmpNodes, &QueueNode{node.Node.Right, node.Pos*2 + 1, level})
				rightPos = node.Pos
			}
		}

		tmpMax = (rightPos - leftPos) + 1

		if tmpMax > max {
			max = tmpMax
		}

		parentNodes = tmpNodes
	}

	return max
}
