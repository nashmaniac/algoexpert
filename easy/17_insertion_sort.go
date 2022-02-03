package easy

func InsertionSort(array []int) []int {

	placeToInsert := 0

	for placeToInsert < len(array) {
		currentMin := array[placeToInsert]
		index := placeToInsert
		minIndex := placeToInsert
		for index < len(array) {
			if currentMin > array[index] {
				currentMin = array[index]
				minIndex = index
			}
			index += 1
		}

		array[minIndex], array[placeToInsert] = array[placeToInsert], array[minIndex]
		placeToInsert += 1
	}
	return array
}
