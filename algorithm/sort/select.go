package sort

func SelectSort(data []int64, reverse bool) []int64 {
	length := len(data)
	if length == 0 {
		return data
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			var condition bool
			if reverse {
				condition = data[j] > data[i]
			} else {
				condition = data[j] < data[i]
			}
			if condition {
				// 交换位置
				data[j], data[i] = data[i], data[j]
			}
		}
	}
	return data
}
