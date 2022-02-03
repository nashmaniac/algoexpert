package easy

func BinarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := (left + right) / 2
		if target == array[mid] {
			return mid
		} else if target < array[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// Write your code here.
	return -1
}
