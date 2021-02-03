package sort

func InsertSortFirst(data []int64) []int64 {
	length := len(data)
	for i := 1; i < length; i++ {
		current := data[i]
		for j := i - 1; j >= 0; j-- {
			if current > data[j] {
				data[j+1] = current
				break
			} else if j == 0 {
				data[j+1] = data[j]
				data[j] = current
			} else {
				data[j+1] = data[j]
			}
		}
	}
	return data
}

func InsertSortSecond(data []int64) []int64 {
	length := len(data)
	for i := 1; i < length; i++ {
		current := data[i]
		j := i - 1
		for j >= 0 && data[j] > current {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = current
	}
	return data
}
