package problems

import (
	"fmt"
	"math"
)

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

	optimal := make([]int, 5)
	optimal[1] = 1
	optimal[2] = 2
	optimal[3] = 3
	optimal[4] = 4
	for i := 5; i <= length; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			fmt.Println("j:", j, "s:", i-j, "i:", i)
			current := optimal[j] * optimal[i-j]
			if current > max {
				max = current
			}
		}
		optimal = append(optimal, max)
	}
	fmt.Println(optimal)
	return optimal[length]
}

// 3(n-3) > n
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
	l := length % 3

	if l == 0 {
		// 被三整除
		return int(math.Pow(float64(3), float64(length/3)))
	}

	if l == 1 {
		// 能剩下4
		return int(math.Pow(float64(3), float64(length/3-1))) * 4
	}

	if l == 2 {
		// 剩下4
		return int(math.Pow(float64(3), float64(length/3))) * 2
	}
	return 0
}
