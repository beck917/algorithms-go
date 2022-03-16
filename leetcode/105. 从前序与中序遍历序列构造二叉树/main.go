package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}

	var i int
	// 通过中序遍历确定中间节点的位置, 找到root节点后, 中序遍历和前序遍历,左右子树的长度都一样
	// 这样就可以确定前序遍历左半边和右半边
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	// 将左半边传入, 因为slice的边界问题,所以这里i+1
	// 左边是去除root后的,左半边, 右边是去除root后的右半边
	root.Left = buildTree(preorder[1:i+1], inorder[0:i+1])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	// 这里可以返回root, 如果root没有左右子树,那么会跳出递归
	// 这样上层就可以拿到left和right, 继续递归了
	return root
}
