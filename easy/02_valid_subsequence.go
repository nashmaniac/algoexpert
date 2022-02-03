package easy

func IsValidSubsequence(array []int, sequence []int) bool {

	firstIndex := 0
	secondIndex := 0

	for secondIndex < len(sequence) && firstIndex < len(array) {
		if sequence[secondIndex] == array[firstIndex] {
			secondIndex++
		}
		firstIndex++
	}

	return secondIndex == len(sequence)
}
