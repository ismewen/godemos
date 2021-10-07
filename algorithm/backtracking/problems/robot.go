/*
m * n 的方格. 从0，0 出发
*/
package problems

import (
	"fmt"
	"strconv"
)

const (
	ROBOTCOW = 3
	ROBOTCOl = 1
	K        = 0
)

func Robot() {
	chess := [ROBOTCOW][ROBOTCOl]int{}
	visited := [ROBOTCOW * ROBOTCOl]int{}

	Visit(&chess, &visited, 0, 0)
	sum := 0
	for _, i := range visited {
		sum += i
	}
	fmt.Printf("M: %d, N: %d, 可进入的格子数量为: %d\n", ROBOTCOW, ROBOTCOl, sum)
}

func GetSum(i int) int {
	str := fmt.Sprintf("%d", i)
	sum := 0
	for j := 0; j < len(str); j++ {
		item, _ := strconv.Atoi(string(str[j]))
		sum += item
	}
	return sum
}

func GetSumTwo(i int) int {
	sum := 0
	for i != 0 {
		sum += i % 10
		i = i / 10
	}
	return sum
}

func Visit(chess *[ROBOTCOW][ROBOTCOl]int, visited *[ROBOTCOW * ROBOTCOl]int, cow, col int) {
	if cow >= 0 && cow < ROBOTCOW && col >= 0 && col < ROBOTCOl && GetSum(cow)+GetSum(col) <= K && visited[cow*ROBOTCOW+col] != 1 {
		// 可以访问
		visited[cow*ROBOTCOW+col] = 1
		Visit(chess, visited, cow+1, col)
		Visit(chess, visited, cow-1, col)
		Visit(chess, visited, cow, col+1)
		Visit(chess, visited, cow, col-1)
	}
}
