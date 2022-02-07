package medium

func SpiralTraverse(array [][]int) []int {
	if len(array) == 0 {
		return nil
	}

	arr := make([]int, 0)
	startRow := 0
	endRow := len(array) - 1
	startCol := 0
	endCol := len(array[0]) - 1

	for startRow <= endRow && startCol <= endCol {
		// traverse horizontal
		for i := startCol; i < endCol+1; i++ {
			arr = append(arr, array[startRow][i])
		}

		for i := startRow + 1; i < endRow+1; i++ {
			arr = append(arr, array[i][endCol])
		}

		for i := endCol - 1; i >= startCol; i-- {
			if startRow == endRow {
				break
			}
			arr = append(arr, array[endRow][i])
		}

		for i := endRow - 1; i > startRow; i-- {
			if startCol == endCol {
				break
			}
			arr = append(arr, array[i][startCol])
		}

		startRow++
		endRow--
		startCol++
		endCol--
	}
	return arr
}
