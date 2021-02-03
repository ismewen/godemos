package search

func binarySearchRecursion(arr []int, value, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == value {
		return mid
	}
	if arr[mid] > value {
		return binarySearchRecursion(arr, value, left, mid)
	} else {
		return binarySearchRecursion(arr, value, mid+1, right)
	}
}

func BinarySearchRecursion(arr []int, value int) int {
	left := 0
	right := len(arr) - 1
	return binarySearchRecursion(arr, value, left, right)
}

func BinarySearchIter(arr []int, value int) int {
	left := 0
	right := len(arr) - 1
	mid := (right-left)/2 + left
	for left <= right {

		if arr[mid] == value {
			return mid
		} else if arr[mid] > value {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (right-left)/2 + left
	}
	return -1
}
