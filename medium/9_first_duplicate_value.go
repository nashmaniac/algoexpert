package medium

type FirstDuplicateValueProblem struct {
}

func FirstDuplicateValue(array []int) int {
	f := FirstDuplicateValueProblem{}
	return f.Variation1(array)
}

func (f *FirstDuplicateValueProblem) Variation1(array []int) int {
	var hashMap map[int]bool = make(map[int]bool)

	for i := 0; i < len(array); i++ {
		if _, ok := hashMap[array[i]]; !ok {
			hashMap[array[i]] = true
		} else {
			return array[i]
		}
	}
	return -1
}
