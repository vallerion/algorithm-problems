package _1684_Count_the_Number_of_Consistent_Strings

func countConsistentStrings(allowed string, words []string) int {
	mp := make(map[rune]bool, len(allowed))
	consistents := 0

	for _, s := range allowed {
		mp[s] = true
	}

	for i := 0; i < len(words); i++ {
		consistents++
		for _, s := range words[i] {
			if _, ok := mp[s]; ok == false {
				consistents--
				break
			}
		}
	}

	return consistents
}
