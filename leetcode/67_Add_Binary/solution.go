package _67_Add_Binary

// "1010"
// "1011"

func addBinary(a string, b string) string {
	smallest, largest := []rune(a), []rune(b)

	if len(smallest) > len(largest) {
		temp := largest
		largest = smallest
		smallest = temp
	}

	i := len(largest) - 1

	diffLength := len(largest) - len(smallest)
	d := 0

	for i >= 0 {
		sum := rune(0)
		if i-diffLength >= 0 {
			sum = (largest[i] - '0') + (smallest[i-diffLength] - '0') + rune(d) // 0, 1, 2
		} else {
			sum = (largest[i] - '0') + rune(d)
		}
		d = 0

		if sum == 2 {
			sum = 0
			d = 1
		}

		if sum == 3 {
			sum = 1
			d = 1
		}

		largest[i] = sum + '0'
		i--
	}

	if d > 0 {
		largest = append([]rune{'1'}, largest...)
	}

	return string(largest)
}
