package main

import "algorithm/backtracking/problems"

func Test(r *[][8]int, a [8]int){
	a[1] = 0
	*r = append(*r, a)
}
func main() {
	//res := make([][8]int, 0)
	//b := [8]int{1, 2, 3}
	//Test(&res, b)
	//fmt.Println(b)
	//fmt.Println(res)
	//problems.EightQueue()
	//print(problems.Matrix())

	print(problems.GetSumTwo(34))
}
