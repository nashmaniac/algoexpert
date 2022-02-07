package medium

import (
	"sort"
)

type MergeOverlappingProblem struct {
}

type Interval [][]int

func (interval Interval) Len() int {
	return len(interval)
}

func (interval Interval) Less(i, j int) bool {
	return interval[i][0] < interval[j][0]
}

func (interval Interval) Swap(i, j int) {
	interval[i], interval[j] = interval[j], interval[i]
}

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	m := MergeOverlappingProblem{}
	return m.Variation1(intervals)
}

func (m *MergeOverlappingProblem) Variation1(intervals [][]int) [][]int {
	sort.Sort(Interval(intervals))
	arr := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		if len(arr) == 0 {
			arr = append(arr, intervals[i])
		} else {
			if arr[len(arr)-1][1] >= intervals[i][0] {
				if intervals[i][1] >= arr[len(arr)-1][1] {
					arr[len(arr)-1][1] = intervals[i][1]
				}
			} else {
				arr = append(arr, intervals[i])
			}
		}
	}
	return arr
}
