package medium

type GroupAnagramProblem struct {
}

func calculateArrayFromWord(str string) []int {
	arr := make([]int, 26)
	for i := 0; i < 26; i++ {
		arr[i] = 0
	}

	for i := 0; i < len(str); i++ {
		arr[int(str[i])-97] += 1
	}

	return arr
}

func (g *GroupAnagramProblem) Variation1(words []string) [][]string {

	visited := make([]bool, len(words))
	wordList := make([][]string, 0)
	for i := 0; i < len(words); i++ {
		if visited[i] {
			continue
		}
		// init the word list
		visited[i] = true
		arr := make([]string, 0)
		arr = append(arr, words[i])
		refArr := calculateArrayFromWord(words[i])

		for j := i + 1; j < len(words); j++ {
			tempArr := calculateArrayFromWord(words[j])
			isMatch := true
			for k := 0; k < 26; k++ {
				if tempArr[k] != refArr[k] {
					isMatch = false
					break
				}
			}
			if isMatch {
				arr = append(arr, words[j])
				visited[j] = true
			}
		}

		wordList = append(wordList, arr)
	}
	return wordList
}

func GroupAnagrams(words []string) [][]string {
	g := GroupAnagramProblem{}
	return g.Variation1(words)
}
