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

// 前序遍历  根左右
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

// 中序遍历  左根右
func (node *TreeNode) InOrder() {
	if node == nil {
		return
	}
	// 中序遍历 左-->根-->右
	// 对左子树进行中序遍历
	node.Left.InOrder()
	// 输出根
	fmt.Printf("%v\t", node.Val)
	// 中序遍历右子树
	node.Right.InOrder()
}

// 后序遍历 左右根
func (node *TreeNode) PostOrder() {
	if node == nil {
		return
	}
	// 后序遍历 左-->右-->根
	// 对左子树后序遍历
	node.Left.PostOrder()
	// 对右子树后序遍历
	node.Right.PostOrder()
	// 输出根
	fmt.Printf("%v\t", node.Val)
}

func CreateNode(val int, left, right *TreeNode) *TreeNode {
	node := &TreeNode{
		Val: val,
		Left: left,
		Right: right,
	}
	return node
}