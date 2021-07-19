package _1684_Count_the_Number_of_Consistent_Strings

func countConsistentStrings(allowed string, words []string) int {
	mp := make(map[rune]bool)

	ar := []rune(allowed)
	for i := 0; i < len(ar); i++ {
		mp[ar[i]] = true
	}

	counter := 0
	for i := 0; i < len(words); i++ {
		word := []rune(words[i])

		for j := 0; j < len(word); j++ {
			_, ok := mp[word[j]]
			if ok == false {
				break
			}
			if j == len(word)-1 {
				counter++
			}
		}
	}

	return counter
}
