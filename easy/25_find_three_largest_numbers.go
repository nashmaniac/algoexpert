package easy

import "math"

func shiftAndUpdate(array []int, el int, index int) {

	for i := 0; i < index+1; i++ {
		if i == index {
			array[i] = el
		} else {
			array[i] = array[i+1]
		}
	}
}
func updateLargest(array []int, el int) {
	if el > array[2] {
		shiftAndUpdate(array, el, 2)
	} else if el > array[1] {
		shiftAndUpdate(array, el, 1)
	} else if el > array[0] {
		shiftAndUpdate(array, el, 0)
	}
}

func FindThreeLargestNumbers(array []int) []int {
	var arr []int = []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, el := range array {
		updateLargest(arr, el)
	}
	return arr
}
