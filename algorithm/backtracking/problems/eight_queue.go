package problems

import (
	"fmt"
	"math"
)

// 8 * 8 的棋盘， 摆放八个皇后，同行、同列、斜对线不能交
func EightQueue() {
	// 存放结果
	queses := make([][8]int, 0)
	queue := [8]int{}

	FindQueue(&queses, queue, 0)

	fmt.Println(len(queses))
}

func FindQueue(queues *[][8]int, queue [8]int, currentCow int) {
	if currentCow >= 8 {
		// 第九行, 说明八皇后已归位, 结束
		fmt.Println("归位", queue, currentCow)
		*queues = append(*queues, queue)
		return
	}
	for col := 0; col < 8; col++ {
		flag := HasConflict(queue, currentCow, col)
		if !flag {
			// 没有冲突，开始摆放下一个
			queue[currentCow] = col
			FindQueue(queues, queue, currentCow+1)
		}
		// 有冲突， 在本行尝试下一个位置
	}

}

func HasConflict(queue [8]int, currentCow, currentCol int) bool {
	for i := 0; i < currentCow; i++ {
		// 获取 当前摆放位置 和 历史摆放位置 的绝对值
		// 历史摆放位置
		history := queue[i]
		m := int(math.Abs(float64(currentCol - history)))
		if m == 0 {
			// 在同一行
			return true
		}
		if m == currentCow-i {
			// 在 对角线
			return true
		}
	}
	//没有任何冲突
	return false
}
