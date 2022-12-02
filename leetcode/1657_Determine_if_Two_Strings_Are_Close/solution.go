package _1657_Determine_if_Two_Strings_Are_Close

import "sort"

// "cabbba", "abbccc"
// cabbba -> caabbb
// caabbb -> baaccc
// baaccc -> abbccc

// c -> 1, a -> 2, b -> 3
// a -> 1, b -> 2, c -> 3

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1, freq2 := make([]int, 26), make([]int, 26)

	for i := 0; i < len(word1); i++ {
		freq1[word1[i]-'a']++
		freq2[word2[i]-'a']++
	}

	for i := 0; i < len(freq1); i++ {
		if freq1[i] > 0 && freq2[i] == 0 {
			return false
		}
		if freq2[i] > 0 && freq1[i] == 0 {
			return false
		}
	}

	sort.Ints(freq1)
	sort.Ints(freq2)

	for i := 0; i < len(freq1); i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}
