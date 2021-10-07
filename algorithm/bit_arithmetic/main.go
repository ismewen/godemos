package main

import (
	"algorithm/bit_arithmetic/problems"
	"fmt"
)

func main() {
	i := problems.NumberOfOneMethodOne(4)
	fmt.Println(i)
	i = problems.NumberOfOneMethodTwo(4)
	fmt.Println(i)
}
