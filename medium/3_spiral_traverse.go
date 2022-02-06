package medium

import "log"

func SpiralTraverse(array [][]int) []int {
	if len(array) == 0 {
		return nil
	}

	arr := make([]int, 0)
	m := len(array)
	n := len(array[0])
	i, j := 0, 0

	for i < m/2 {
		// traverse horizontal
		for j < n {
			arr = append(arr, array[i][j])
			log.Println(arr)
			j++
		}
		j--
		i++
		for i < m {
			arr = append(arr, array[i][j])
			i++
		}
		i--
		j--
		for j >= 0 {
			arr = append(arr, array[i][j])
			j--
		}
		j++
		i--
		for i > m {
			arr = append(arr, array[i][j])
			i--
		}
		m++
		n--
	}
	return arr
}
