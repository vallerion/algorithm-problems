package _415_Add_Strings

func addStrings(num1 string, num2 string) string {
	num1r, num2r := []rune(num1), []rune(num2)

	bigger, smaller := num1r, num2r

	if len(bigger) < len(smaller) {
		bigger, smaller = smaller, bigger
	}

	i := len(bigger) - 1
	difLength := len(bigger) - len(smaller)
	d := 0

	for i >= 0 {
		sum := 0
		if i-difLength >= 0 {
			sum = int(bigger[i]-'0') + int(smaller[i-difLength]-'0') + d
		} else {
			sum = int(bigger[i]-'0') + d
		}
		d = 0

		if sum >= 10 {
			sum %= 10
			d = 1
		}

		bigger[i] = rune(sum) + '0'
		i--
	}

	if d > 0 {
		bigger = append([]rune{'1'}, bigger...)
	}

	return string(bigger)
}
