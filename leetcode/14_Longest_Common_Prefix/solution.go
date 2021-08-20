package _14_Longest_Common_Prefix

// ["flower","flow","flight"]

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	i := 0

	for i < len(strs[0]) {
		sym := strs[0][i]

		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != sym {
				return strs[0][:i]
			}
		}

		i++
	}

	if i == 0 {
		return ""
	}

	return strs[0][:i]
}
