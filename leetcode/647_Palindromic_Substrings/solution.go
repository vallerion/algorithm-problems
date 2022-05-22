package _647_Palindromic_Substrings

// aaaa
// a
// a -> aaa
// a -> aaa
// a

func countSubstrings(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			res++
			left--
			right++
		}

		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			res++
			left--
			right++
		}
	}

	return res
}

func countSubstringsOn3(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				cnt++
			}
		}
	}

	return cnt + len(s)
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
