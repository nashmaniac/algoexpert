package medium

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	arr := make([][]int, 0)
	for i := 0; i < len(array)-2; i++ {
		leftIndex := i + 1
		rightIndex := len(array) - 1
		for leftIndex < rightIndex {
			leftNum := array[leftIndex]
			rightNum := array[rightIndex]
			sum := array[i] + leftNum + rightNum
			if sum < target {
				leftIndex++
			} else if sum > target {
				rightIndex--
			} else {
				arr = append(arr, []int{array[i], leftNum, rightNum})
				leftIndex++
				rightIndex--
			}
		}

	}
	return arr
}
