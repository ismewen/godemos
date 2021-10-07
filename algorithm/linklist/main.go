package main

import "algorithm/linklist/types"

func Reverse(head *types.Head,node *types.Node) *types.Node{
	if node.Next != nil {
		next := Reverse(head, node.Next)
		next.Next = node
	}else{
		firstNode := head.Next
		firstNode.Next = nil
		head.Next = node
	}
	return node
}

func main(){
	head := types.NewHead("reverse")
	for i:=0; i<10; i++ {
		head.Insert(i, &types.Node{Data: i})
	}
	head.Print()
	Reverse(head, head.Next)
	head.Print()
}