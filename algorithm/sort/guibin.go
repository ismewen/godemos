package sort

func Merge(left, right []int) []int {
	// 将data[i]
	length := len(left) + len(right)
	i := 0
	j := 0
	temp := make([]int, length)
	for k := 0; k < length; k++ {
		var thisOne int
		if i > len(left)-1 {
			thisOne = right[j]
			j++
		} else if j > len(right)-1 {
			thisOne = left[i]
			i++
		} else if left[i] < right[j] {
			thisOne = left[i]
			i++
		} else {
			thisOne = right[j]
			j++
		}
		temp[k] = thisOne
	}
	return temp
}

func GuibinSort(data []int) []int {
	length := len(data)
	if length == 1{
		return data
	}
	mid := length / 2
	left := data[0:mid]
	right := data[mid:length]
	// 左边
	gLeft := GuibinSort(left)
	gRight := GuibinSort(right)
	data = Merge(gLeft, gRight)
	return data
}

