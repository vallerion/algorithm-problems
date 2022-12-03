package _1704_Determine_if_String_Halves_Are_Alike

// valnet -> a e -> 1 1 -> true
// aaanet -> 3 1 -> false

var vowels = map[rune]int{
	'a': 0,
	'e': 0,
	'i': 0,
	'o': 0,
	'u': 0,
	'A': 0,
	'E': 0,
	'I': 0,
	'O': 0,
	'U': 0,
}

func halvesAreAlike(s string) bool {
	counter1, counter2 := 0, 0

	for i, v := range s {
		if _, ok := vowels[v]; ok {
			if i < len(s)/2 {
				counter1++
			} else {
				counter2++
			}
		}
	}

	return counter2 == counter1
}
