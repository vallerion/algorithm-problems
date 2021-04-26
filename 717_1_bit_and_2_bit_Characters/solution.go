package _717_1_bit_and_2_bit_Characters

// 0 1 1 1 0
// 0 1 0
// 0 0 1 0
// 1 1 1 0
// 1 1 0

// 1 1 -> 2
// 0
func isOneBitCharacter(bits []int) bool {
	currSum := 0
	i := 0

	for i < len(bits) {
		currSum = 0
		if bits[i] == 0 {
			i++
			continue
		}

		if bits[i] == 1 && i+1 < len(bits) {
			currSum = bits[i] + bits[i+1]
		} else {
			currSum = bits[i]
		}

		i += 2
	}

	return currSum == 0
}
