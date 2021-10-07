package problems

// 统计 二进制 1的个数

func NumberOfOneMethodOne(i int) int {
	flag := 1
	count := 0
	for flag != 0 {
		if flag&i != 0 {
			count++
		}
		flag = flag << 1
	}
	return count
}

func NumberOfOneMethodTwo(i int) int {
	count := 0
	for i != 0 {
		i = i & (i - 1)
		count++
	}
	return count
}
