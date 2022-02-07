package medium

type LongestPeakProblem struct {
}

func LongestPeak(array []int) int {
	l := LongestPeakProblem{}
	return l.Variation1(array)
}

func (l *LongestPeakProblem) Variation1(array []int) int {
	longestPeakLength := 0
	i := 1
	for i < len(array)-1 {
		isPeak := array[i-1] < array[i] && array[i] < array[i+1]
		if !isPeak {
			i += 1
			continue
		}

		leftIndex := i - 2
		for leftIndex >= 0 && array[leftIndex] < array[leftIndex+1] {
			leftIndex -= 1
		}

		rightIndex := i + 2
		for rightIndex < len(array) && array[rightIndex] < array[rightIndex-1] {
			rightIndex += 1	
		}
		currentLength := rightIndex - leftIndex - 1
		if currentLength > longestPeakLength {
			longestPeakLength = currentLength
		}
		i = rightIndex
	}
	return longestPeakLength
}
