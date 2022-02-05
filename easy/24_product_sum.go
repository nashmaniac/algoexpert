package easy

type SpecialArray []interface{}

func helper(arr SpecialArray, multiplier int) int {
	sum := 0
	for _, el := range arr {
		if cast, ok := el.(SpecialArray); !ok {
			sum += helper(cast, multiplier+1)
		}
		if cast, ok := el.(int); !ok {
			sum += cast
		}
	}
	return sum * multiplier
}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	return helper(array, 1)
}
