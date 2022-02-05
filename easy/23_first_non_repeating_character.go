package easy

func FirstNonRepeatingCharacter(str string) int {
	var charMap map[rune]int = make(map[rune]int)

	for _, ch := range str {
		charMap[ch] += 1
	}
	for idx, ch := range str {
		if charMap[ch] == 1 {
			return idx
		}
	}
	return -1
}
