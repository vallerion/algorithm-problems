package _3_Longest_Substring_Without_Repeating_Characters

import "fmt"

func lengthOfLongestSubstring(s string) int {
	longestSub, currentSub := 0, 0
	wordMap := make(map[rune]int, 0)
	ss := []rune(s)

	i, j := 0, 0

	for i < len(ss) && j < len(ss) {
		j = i

		if _, exists := wordMap[ss[j]]; exists {
			fmt.Println(wordMap)
			i = j + 1
			wordMap = make(map[rune]int, 0)
			currentSub = 0
			continue
		}

		currentSub++
		wordMap[ss[j]] = 0
		longestSub = max(longestSub, currentSub)

		j++
	}

	return longestSub
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
