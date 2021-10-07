package problems

// 统计 二进制 1的个数

func NumInt(i int) {
	flag := 1
	count := 0
	for i&flag != 0 {
		count++
		flag = 1 << flag
	}
}
