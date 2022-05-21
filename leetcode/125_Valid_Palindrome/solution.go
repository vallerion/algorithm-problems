package _125_Valid_Palindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	replaced := ""
	for i := 0; i < len(s); i++ {
		if (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') {
			replaced += string(s[i])
		}
	}

	i, j := 0, len(replaced)-1

	for i < j {
		if replaced[i] != replaced[j] {
			return false
		}
		i++
		j--
	}

	return true
}
