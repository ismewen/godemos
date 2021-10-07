package problems

import "fmt"

// 判断矩阵中是否存在某路径

const (
	COM = 3
	COL = 4
)

func Matrix() bool {
	chess := [COM][COL]byte{
		{'a', 'b', 'f', 'b'},
		{'c', 'f', 'c', 'f'},
		{'j', 'g', 'f', 'c'},
	}
	path := []byte{'b', 'f', 'c', 'f', 'g'}
	visited := &[COM][COL]bool{}
	for cow := 0; cow < COM; cow++ {
		for col := 0; col < COL; col++ {
			fmt.Println(chess[cow][col])
			if HasMatched(&chess, visited, cow, col, path, 0) {
				return true
			}
		}
	}
	return false
}

func HasMatched(chess *[COM][COL]byte, visited *[COM][COL]bool, cow, col int, path []byte, matched int) bool {
	fmt.Println(visited)
	if matched == len(path) {
		return true
	}
	flag := false
	if cow >= 0 && cow < COM && col >= 0 && col < COL && !visited[cow][col] {
		visited[cow][col] = true
		if chess[cow][col] == path[matched] {
			matched += 1
			flag =
				HasMatched(chess, visited, cow+1, col, path, matched) ||
					HasMatched(chess, visited, cow-1, col, path, matched) ||
					HasMatched(chess, visited, cow, col+1, path, matched) ||
					HasMatched(chess, visited, cow, col-1, path, matched)
		}

	}
	if !flag{
		// 没找到，抹除访问痕迹
		if cow >= 0 && cow < COM && col >=0 && col < COL {
			visited[cow][col] = false
		}
	}
	return flag
}
