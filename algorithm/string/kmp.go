package string

func GetNext1(t string) []int {
	length := len(t)
	next := make([]int, length)
	j := 0
	next[0] = 0
	for i := 1; i < length; {
		// 错开对比
		if j == 0 && t[i] != t[j] {
			// 回溯到0
			next[i] = 0
			i++
			continue
		}
		if t[j] == t[i] {
			// 存在对称
			j++
			next[i] = j
			i++
		} else {
			// 回溯
			j = next[j-1]
		}
	}
	return next
}

func Kmp1(m string, t string) bool {
	next := GetNext(t)
	j := 0
	i := 0
	for ; i < len(m) && j < len(t); {
		if m[i] == t[j] {
			// 对比前行
			i++
			j++
		} else if j > 0 {
			// 可以回溯
			j = next[j-1]
		} else {
			// 回溯到0, 且不等，后移i
			i++
		}
	}
	return j == len(t)
}

func GetNext(s string) []int {
	next := make([]int, len(s))
	next[0] = 0
	i := 1
	for i < len(s) {
		j := 0
		for j < len(s) {
			if s[i] == s[j] {
				j++
				next[i]  = j
				i++
			}else if j == 0 {
				next[i]	= 0
				i++
				break
			}else{
				j = next[j-1]
			}
		}
	}
	return next
}

func Kmp(m string, s string) bool {
	next := GetNext(s)
	i := 0
	j := 0
	for i < len(m) && j < len(s) {
		if m[i] == m[j] {
			i++
			j++
		} else if j > 0 {
			j = next[j-1]
		} else {
			i++
		}
	}
	return j == len(s)
}
