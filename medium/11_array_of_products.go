package medium

type ArrayOfProductsProblem struct {
}

func (a *ArrayOfProductsProblem) Variation2(arr []int) []int {
	left := make([]int, len(arr))
	right := make([]int, len(arr))
	// intialize array
	leftProduct := 1
	rightProduct := 1

	for i := 0; i < len(arr); i++ {
		left[i] = leftProduct
		right[len(arr)-1-i] = rightProduct

		leftProduct *= arr[i]
		rightProduct *= arr[len(arr)-1-i]
	}

	for i := 0; i < len(left); i++ {
		left[i] = left[i]*right[i]
	}
	return left
}
func (a *ArrayOfProductsProblem) Variation1(arr []int) []int {
	array := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		product := 1
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}

			product *= arr[j]
		}
		array[i] = product
	}
	return array
}

func ArrayOfProducts(array []int) []int {
	a := ArrayOfProductsProblem{}
	return a.Variation1(array)
}
