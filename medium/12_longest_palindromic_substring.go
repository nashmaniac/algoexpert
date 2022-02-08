package medium

type LongestPalindromicSubstringProblem struct {
}

func LongestPalindromicSubstring(str string) string {
	l := LongestPalindromicSubstringProblem{}
	return l.Variation1(str)
}

func findPalindrome(str string, startIndex int, endIndex int) []int {
	for startIndex >= 0 && endIndex < len(str) {
		if str[startIndex] != str[endIndex] {
			break
		}
		startIndex -= 1
		endIndex += 1
	}
	return []int{startIndex + 1, endIndex}
}

func (l *LongestPalindromicSubstringProblem) Variation1(str string) string {
	if str == "" || len(str) == 0 {
		return ""
	}
	arr := []int{0, 1}
	count := 1

	for i := 1; i < len(str); i++ {
		left := findPalindrome(str, i-1, i+1)
		right := findPalindrome(str, i-1, i)

		leftLength := left[1] - left[0]
		rightLength := right[1] - right[0]

		if leftLength > rightLength {
			if leftLength > count {
				count = leftLength
				arr = left
			}
		} else if rightLength > leftLength {
			if rightLength > count {
				count = rightLength
				arr = right
			}
		} else {
		}
	}
	return str[arr[0]:arr[1]]
}
