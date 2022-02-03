package easy

func InsertionSort(array []int) []int {

	for i := 0;i<len(array); i++ {
		for j := i;j>0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}
