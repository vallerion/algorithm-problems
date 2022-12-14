package _242_Valid_Anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
		freq[t[i]]--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}

	return true
}