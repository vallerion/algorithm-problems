package _28_Implement_strStr

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		if len(needle) == 0 {
			return 0
		} else {
			return -1
		}
	}

	if len(needle) == 0 {
		return 0
	}

	haystackR, needleR := []rune(haystack), []rune(needle)

	hashNeedle := 0
	for i := 0; i < len(needleR); i++ {
		hashNeedle += int(needleR[i])
	}

	hashHaystack := 0

	for i := 0; i < len(haystackR); i++ {
		hashHaystack += int(haystackR[i])

		if i >= len(needleR)-1 {
			if hashHaystack == hashNeedle && deepCompare(haystackR, needleR, i-(len(needleR)-1)) {
				return i-(len(needleR)-1)
			}
			hashHaystack -= int(haystackR[i-(len(needleR)-1)])
		}
	}

	return -1
}

func deepCompare(haystack, needle []rune, start int) bool {
	for i := start; i < start+len(needle); i++ {
		if haystack[i] != needle[i-start] {
			return false
		}
	}
	return true
}
