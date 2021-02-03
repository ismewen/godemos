package sort

func ShellSort(data []int) []int {
	gap := 1
	n := 3
	length := len(data)
	for gap < length/n {
		gap = gap*n + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := data[i]
			j := i - gap
			for j >= 0 && data[j] > temp {
				// 往后偏移
				data[j+gap] = data[j]
				j = j - gap
			}
			data[j+gap] = temp
		}
		println("what is gap", gap)
		gap = gap / n
	}

	return data
}
