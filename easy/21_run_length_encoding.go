package easy

import "strconv"

func RunLengthEncoding(str string) string {
	var bytes []byte

	count := 1
	var ch byte = str[0]

	for i := 1; i < len(str); i++ {
		if ch != str[i] {
			// break found
			bytes = append(bytes, []byte(strconv.Itoa(count))[0])
			bytes = append(bytes, ch)
			count = 1
			ch = str[i]

		} else {
			if count == 9 {
				bytes = append(bytes, []byte(strconv.Itoa(count))[0])
				bytes = append(bytes, ch)
				count = 1
			} else {
				count++
			}
		}
	}
	bytes = append(bytes, []byte(strconv.Itoa(count))[0])
	bytes = append(bytes, ch)

	return string(bytes)
}
