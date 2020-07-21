package binaryTree

import "fmt"

/*
* 二叉树节点
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
* 二叉树
*/
type BinaryTree struct {
	root *TreeNode
}

func (node *TreeNode) PreOrder() {
	if node == nil {
		return
	}
	// 前序遍历  根-->左-->右
	// 先输出根
	fmt.Printf("%v\t", node.Val)
	// 对左子树进行前序遍历
	node.Left.PreOrder()
	// 对右子树进行前序遍历
	node.Right.PreOrder()
}

func CreateNode(val int, left, right *TreeNode) *TreeNode {
	node := &TreeNode{
		Val: val,
		Left: left,
		Right: right,
	}
	return node
}