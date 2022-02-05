package medium

func MoveElementToEnd(array []int, toMove int) []int {
	indx := 0
	for i, el := range array {
		if el != toMove {
			array[indx], array[i] = array[i], array[indx]
			indx++
		}
	}
	return array
}