package _58_Length_of_Last_Word

func lengthOfLastWord(s string) int {
	sr := []rune(s)
	length := 0
	inWord := false

	for i := len(sr) - 1; i >= 0; i-- {
		if sr[i] == ' ' {
			if inWord == true {
				break
			} else {
				continue
			}
		} else {
			length++
			inWord = true
		}

	}

	return length
}