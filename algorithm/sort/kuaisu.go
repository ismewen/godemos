package sort

func quickSort(arr []int, lo, hi int) []int {
	if lo >= hi {
		return nil
	}
	j := Partition(arr, lo, hi)
	quickSort(arr, lo, j-1)
	quickSort(arr, j+1, hi)
	return arr
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func Partition(arr []int, lo, hi int) int {
	//
	i := lo + 1
	j := hi
	for true {
		for arr[i] < arr[lo] && i < hi {
			i++
		}
		for arr[j] > arr[lo] && j > lo {
			j--
		}

		if j <= i {
			break
		}
		arr[j], arr[i] = arr[i], arr[j]
	}
	arr[lo], arr[j] = arr[j], arr[lo]
	return j
}
