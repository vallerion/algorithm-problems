package _28_Implement_strStr

// "bba"

// j=0,i=1 -> b==b -> [0 1 0],i++,j++
// j=1,i=2 -> b!=a -> j = 0
// j=0,i=2 -> b!=a ->

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if len(needle) == 0 && len(haystack) == 0 {
		return 0
	}

	pi := makePrefix(needle)

	l, k := 0, 0

	for l < len(haystack) && k < len(needle) {
		if haystack[l] == needle[k] {
			l++
			k++
		} else if k == 0 {
			l++
		} else {
			k = pi[k-1]
		}
	}

	if k == len(needle) {
		return l - len(needle)
	}

	return -1
}

func makePrefix(str string) []int {
	st := []rune(str)
	pi := make([]int, len(str))
	j, i := 0, 1

	for i < len(st) {

		if st[i] == st[j] {
			pi[i] = j + 1
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = pi[j-1]
		}
	}

	return pi
}
