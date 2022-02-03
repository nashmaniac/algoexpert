package easy

func CaesarCipherEncryptor(str string, key int) string {
	// Write your code here.
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		runes[i] = rune((((int(str[i]) - 97) + key) % 26) + 97)
	}
	return string(runes)
}
