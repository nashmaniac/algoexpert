package medium

import "math"

func SmallestDifference(array1, array2 []int) []int {

	arr := make([]int, 2)
	smallestCount := math.MaxInt32
	diff := 0

	firstIndex := 0
	secondIndex := 0

	for firstIndex < len(array1) && secondIndex < len(array2) {
		firstNum := array1[firstIndex]
		secondNum := array2[secondIndex]
		if firstNum < secondNum {
			diff = secondNum - firstNum
			firstIndex++
		} else if firstNum > secondNum {
			diff = firstNum - secondNum
			secondIndex++
		} else {
			return []int{firstNum, secondNum}
		}
		if diff < smallestCount {
			smallestCount = diff
			arr[0] = firstNum
			arr[1] = secondNum
		}
	}

	return arr
}
