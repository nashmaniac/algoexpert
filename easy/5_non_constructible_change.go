package easy

import "sort"

func NonConstructibleChange(coins []int) int {

	sort.Ints(coins)
	change := 0
	for i := 0; i < len(coins); i++ {
		if change+1 < coins[i] {
			return change + 1
		} else {
			change += coins[i]
		}
	}
	return change + 1
}
