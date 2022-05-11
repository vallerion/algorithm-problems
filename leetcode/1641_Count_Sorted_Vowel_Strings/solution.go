package _1641_Count_Sorted_Vowel_Strings

// O(n*m), m - vowels (a, e, i, o, u) -> O(5n) -> O(n)
func countVowelStringsO5n(n int) int {
	res := 0
	help(0, 0, &res, n)
	return res
}

func help(start, deep int, res *int, n int) {
	if deep == n {
		*res++
		return
	}

	for i := start; i < 5; i++ {
		help(i+1, deep+1, res, n)
	}
}
