package _6074_Percentage_of_Letter_in_String

// "foobar"
// "o"
// "jjjj"
// "k"
// "vmvvvvvzrvvpvdvvvvyfvdvvvvpkvvbvvkvvfkvvvkvbvvnvvomvzvvvdvvvkvvvvvvvvvlvcvilaqvvhoevvlmvhvkvtgwfvvzy"
// "v"

func percentageLetter(s string, letter byte) int {
	cnt := 0.

	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			cnt++
		}
	}

	return int(cnt / float64(len(s)) * 100)
}
