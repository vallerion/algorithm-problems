package _17_Letter_Combinations_of_a_Phone_Number

// abc def

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mp := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	temp := make([]byte, 0)
	result := make([]string, 0)

	backtracking(digits, 0, mp, temp, &result)

	return result
}

func backtracking(digits string, i int, digitsMap map[byte][]byte, temp []byte, result *[]string) {
	if len(temp) == len(digits) {
		*result = append(*result, string(temp))
		return
	}

	letters := digitsMap[digits[i]]

	for j := 0; j < len(letters); j++ {
		temp = append(temp, letters[j])
		backtracking(digits, i+1, digitsMap, temp, result)
		temp = temp[:len(temp)-1]
	}
}
