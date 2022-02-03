package easy

import "sort"

func TwoNumberSum(array []int, target int) []int {
	// Solution O(n) time O(n) space
	var hashmap map[int]int = make(map[int]int)

	for i := 0; i < len(array)-1; i++ {
		diff := target - array[i]
		if index, ok := hashmap[array[i]]; !ok {
			hashmap[diff] = array[i]
		} else {
			return []int{array[i], index}
		}
	}
	return []int{}
}

func TwoNumberSumSortingMethod(array []int, target int) []int {
	// solution O(nlogn) time O(1) space
	sort.Ints(array)
	left := 0
	right := len(array) - 1

	for left < right {
		total := array[left] + array[right]
		if total == target {
			return []int{array[left], array[right]}
		} else if total < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}
