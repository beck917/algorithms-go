package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CTreeNode struct {
	node *TreeNode
	code int
}

// 事实上，我们只需要检查最后一个编号是否正确，因为最后一个编号的值最大。
func isCompleteTree(root *TreeNode) bool {
	parentNodes := []*CTreeNode{
		&CTreeNode{root, 1},
	}

	nodeLen := 0
	var lastNode *CTreeNode
	for len(parentNodes) != 0 {
		var tmpNodes []*CTreeNode
		for _, node := range parentNodes {
			if node.node.Left != nil {
				tmpNodes = append(tmpNodes, &CTreeNode{
					node.node.Left, node.code * 2,
				})
			}

			if node.node.Right != nil {
				tmpNodes = append(tmpNodes, &CTreeNode{
					node.node.Right, node.code*2 + 1,
				})
			}
			nodeLen++
			if len(tmpNodes) == 0 {
				lastNode = node
			}
		}
		parentNodes = tmpNodes
	}

	return nodeLen == lastNode.code
}
