package main

import (
	"algorithm/tree/types"
	"fmt"
)

func main() {
	node := types.NewTestTree()
	fmt.Println("\n递归前序遍历")
	res := make([]string, 0)
	types.RecursionPreOrder(node, &res)

	fmt.Println("\n递归前序遍历 From Res:")
	for _, v := range res{
		fmt.Print(v)
	}
	fmt.Println("\n递归中序遍历")
	types.RecursionInOrder(node)
	fmt.Println("\n递归后序遍历")

	types.RecursionPostOrder(node)
	fmt.Println("\n层序遍历")
	types.SequenceOrder(node)
	fmt.Print("\n循环前序遍历")
	types.LoopPreOrder(node)
	fmt.Print("\n循环中序遍历")
	types.LoopInOrder(node)
	fmt.Print("\n循环后序遍历")
	types.LoopPostOrder(node)
	preOrder := []string{"1", "2", "4", "7", "3", "5", "6", "8"}
	inOrder := []string{"4", "7", "2", "1", "5", "3", "8", "6"}
	root, _ := types.RecreateTree(preOrder, inOrder)
	fmt.Println("\n递归前序遍历")
	//types.RecursionPreOrder(root)
	fmt.Println("\n递归中序遍历")
	types.RecursionInOrder(root)

}
