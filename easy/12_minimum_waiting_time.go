package easy

import (
	"sort"
)

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)

	countSum := 0
	for i := 0; i < len(queries)-1; i++ {
		countSum += (queries[i] * (len(queries) - (i+1)))
	}
	return countSum
}
