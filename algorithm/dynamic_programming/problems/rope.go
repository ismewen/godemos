package dynamic_programming

// 一根绳子肠胃n分成，怎么分乘积最大
func DynamicRope(length int) int {
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}
	if length == 4 {
		return 4
	}

	optimal := make([]int, 4)
	optimal[1] = 1
	optimal[2] = 1
	optimal[3] = 2
	optimal[4] = 4

	for i := 5; i <= length; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			current := optimal[j] * optimal[i-j]
			if current > max {
				max = current
			}
		}
		optimal[i] = max
	}
	return optimal[length]
}

func GreedyRope(length int) int {
	if length == 0 {
		return 1
	}
	if length == 2 {
		return 1
	}

	if length == 3 {
		return 2
	}
	if length == 4 {
		return 4
	}
	// length >= 5
	i := 1
	l := length / 3

	for l > 0 {

	}

	return 0
}
