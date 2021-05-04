package _28_Implement_strStr

// Input: haystack = "hello", needle = "ll"
// Output: 2
// h == l / haystack[0] == needle[0] --
// e == l / haystack[1] == needle[0] --
// l == l / haystack[2] == needle[0] ++
// l == l / haystack[3] == needle[1] ++

// mississippi / mississippi
//

func strStr(haystack string, needle string) int {
	haystackR := []rune(haystack)
	needleR := []rune(needle)

	cmpCounter := 0
	//lastIndex := 0

	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	i := 0
	j := 0

	for i < len(haystackR) {
		if haystackR[i] == needleR[j] {
			j++
			cmpCounter++
		} else if cmpCounter > 0 {
			i = i - cmpCounter
			cmpCounter = 0
			j = 0
		}

		if cmpCounter == len(needleR) {
			return i - cmpCounter + 1
		}

		i++
	}

	return -1
}
