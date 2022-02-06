package medium

func IsMonotonic(array []int) bool {
	if len(array) < 3 {
		return true
	}

	var isGreater *bool = nil
	for i := 1; i < len(array); i++ {
		if array[i] == array[i-1] {
			continue
		}
		isCurrentDiffGreater := array[i] > array[i-1]
		if isGreater == nil {
			isGreater = &isCurrentDiffGreater
		} else {
			if *isGreater != isCurrentDiffGreater {
				return false
			}
		}

	}
	return true
}
