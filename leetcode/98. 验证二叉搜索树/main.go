package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	return dfs(root, math.MaxInt64, math.MinInt64)

}

// 首先搜索二叉树的要求,是左子树都<root, 右子树都>root, 注意不只是左右节点
// 这里我用leftprev和rightprev的命名方式,更好理解
// left就是左边的父节点, 子节点要求是<父节点, 所以要层层把新的父节点传递下去, 满足这样的判断后,
/**
看这样一棵树
     6
	/     \
	4      8
	/ \    / \
	2 5    7  9
第三层,比较好解释
在左子树的2肯定是比4小也比6小,只需要一层对比就行了,也就是我们传递的prevleft,实际上是2的父节点
但是我们看下右子节点5, 他要满足两个条件, 1. 比4要大,这个很容易,也就是对比prevright(因为是右节点)
另外还要满足比6小,因为这个节点,是左子树的一部分,这个时候就需要和prevLeft对比, 而这时其实prevleft没有变化
传递的是6,也就是上两层的节点(看再哪里分叉)
*/
func dfs(root *TreeNode, prevLeft, prevRight int) bool {
	if root == nil {
		return true
	}

	//fmt.Println(root.Val, prevLeft, prevRight)
	if root.Val >= prevLeft || root.Val <= prevRight {
		return false
	}

	return dfs(root.Left, root.Val, prevRight) && dfs(root.Right, prevLeft, root.Val)
}
