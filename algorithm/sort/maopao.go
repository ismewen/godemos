package sort

func MaopaoSort(data []int64) []int64 {
	// 从小到大为正序
	// reverse 为从大到小
	length := len(data)
	if length == 0 {
		return data
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if data[j] > data[j+1] {
				// 位置不对，交换位置
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}
