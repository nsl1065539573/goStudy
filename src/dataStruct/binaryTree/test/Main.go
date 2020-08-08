package main

import "dataStruct/binaryTree"

func main()  {
	node7 := binaryTree.CreateNode(7, nil, nil)
	node6 := binaryTree.CreateNode(6, nil, nil)
	node5 := binaryTree.CreateNode(5, nil, nil)
	node4 := binaryTree.CreateNode(4, nil, nil)
	node3 := binaryTree.CreateNode(3, node6, node7)
	node2 := binaryTree.CreateNode(2, node4, node5)
	node1 := binaryTree.CreateNode(1, node2, node3)
	//node1.PreOrder()
	//node1.InOrder()
	node1.PostOrder()
}

func add(a int, b int, arg ...int) int {    //2个或多个参数
	return 0
}