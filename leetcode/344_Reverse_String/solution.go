package _344_Reverse_String

// abc, len=3
// (len - 1) - 2 = 0 // 2 <-> 0
// (len - 1) - 1 = 1 // 1 <-> 1 end

// testim, len=6
// (len - 1) - i = j // i <-> j
// (len - 1) - 5 = 0 // 5 <-> 0
// (len - 1) - 4 = 1 // 4 <-> 1
// (len - 1) - 3 = 2 // 3 <-> 2
// (len - 1) - 2 = 3 // 2 <-> 3 end
func reverseString(s []byte) {
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		j := (length - 1) - i
		if i <= j {
			break
		}

		s[i], s[j] = s[j], s[i]
	}
}