package main

import "algorithm/stack/types"

func main(){
	stack := types.NewStack()

	stack.Push("a1")
	stack.Push("a2")
	stack.Push("a3")
	stack.Print()

}
