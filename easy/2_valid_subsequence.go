package easy

func IsValidSubsequence(array []int, sequence []int) bool {
	// Solution O(n) time O(n) space
	// prepare a hashmap to keep all the index
	var hashmap map[int]int = make(map[int]int)
	for i, elem := range array {
		hashmap[elem] = i
	}

	var lastKnown int = -1
	for _, elem := range sequence {
		var index int
		var ok bool
		if index, ok = hashmap[elem]; !ok {
			return false
		}
		if lastKnown >= index {
			return false
		}
		lastKnown = index
	}

	return true
}
