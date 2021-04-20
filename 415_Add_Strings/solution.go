package _415_Add_Strings

// "3876620623801494171"
// "6529364523802684779"

// 9 + 1 = 10 -> 10 > 9 -> 10 % 10 -> 0
// r=0, d=1, "0"
// 7 + 7 + 1 = 15 -> 15 > 9 -> 15 % 10 -> 5
// r=5, d=1, "50"
// 7 + 1 + 5 = 13 -> 1 / 3

// "99"
// "09"

// 9 + 9 = 18 -> 18 > 9 -> 18 % 10 -> 8
// r=8, d=1, "8"
// 9 + 0 + 1 = 10 -> 10 > 9 -> 10 % 10 -> 0

func AddStrings(num1 string, num2 string) string {
	num1r := []rune(num1)
	num2r := []rune(num2)
	result := make([]rune, 0)

	sum := 0
	mod := 0
	div := 0

	i := len(num1r) - 1
	j := len(num2r) - 1
	//for i := lend - 1; i >= 0; i-- {
	for i >= 0 || j >= 0 {
		d1 := 0
		d2 := 0

		if i >= 0 {
			d1 = int(num1r[i] - '0')
		}
		if j >= 0 {
			d2 = int(num2r[j] - '0')
		}

		sum = d1 + d2 + div
		mod = sum % 10
		div = sum / 10

		if sum > 9 {
			result = append([]rune{rune(mod) + '0'}, result...)
		} else {
			result = append([]rune{rune(sum) + '0'}, result...)
		}

		i--
		j--

		if i < 0 && j < 0 && div != 0 {
			result = append([]rune{rune(div) + '0'}, result...)
		}
	}

	return string(result)
}