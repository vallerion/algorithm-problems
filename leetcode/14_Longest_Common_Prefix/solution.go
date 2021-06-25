package _14_Longest_Common_Prefix

func longestCommonPrefix(strs []string) string {
	k := 0
	shortest := len(strs[0])

	for i := 0; i < len(strs); i++ {
		shortest = min(shortest, len(strs[i]))
	}

	for k < shortest {
		symbolExists := true


		for i := 0; i < len(strs); i++ {
			if strs[i][k] != strs[0][k] {
				symbolExists = false
				break
			}
		}

		if symbolExists == false {
			break
		}

		k++
	}

	if k == 0 {
		return ""
	}

	return strs[0][:k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
