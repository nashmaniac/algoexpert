package easy

func GenerateDocument(characters string, document string) bool {
	var characterMap map[byte]int = make(map[byte]int)
	var documentMap map[byte]int = make(map[byte]int)

	for i := 0; i < len(characters); i++ {
		count, ok := characterMap[characters[i]]
		if !ok {
			characterMap[characters[i]] = 1
		} else {
			characterMap[characters[i]] = count + 1
		}
	}

	for i := 0; i < len(document); i++ {
		count, ok := documentMap[document[i]]
		if !ok {
			documentMap[document[i]] = 1
		} else {
			documentMap[document[i]] = count + 1
		}
	}

	for key, value := range documentMap {
		index, ok := characterMap[key]
		if !ok {
			return false
		}
		if index < value {
			return false
		}
	}

	return true
}
