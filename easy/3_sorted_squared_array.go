package easy

func SortedSquaredAarray(array []int) []int {
	var arr []int = make([]int, len(array))
	left := 0
	right := len(array) - 1
	placeToInsert := len(array) - 1
	for left <= right {
		if array[left]*array[left] > array[right]*array[right] {
			arr[placeToInsert] = array[left] * array[left]
			left++
		} else {
			arr[placeToInsert] = array[right] * array[right]
			right--
		}
		placeToInsert--
	}
	return arr
}
