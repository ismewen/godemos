package main

import (
	"algorithm/dynamic_programming/problems"
	"fmt"
)

func main(){
	i := problems.DynamicRope(10)
	fmt.Println(i)
	i = problems.GreedyRope(10)
	fmt.Println(i)
}
