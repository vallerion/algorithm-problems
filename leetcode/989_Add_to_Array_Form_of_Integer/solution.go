package _989_Add_to_Array_Form_of_Integer

// [9,9], 9
func addToArrayForm(num []int, k int) []int {
	i := len(num) - 1

	d := 0
	for i >= 0 || k > 0 {
		mod := k % 10 // 9 // 0

		sum := 0
		if i >= 0 {
			sum = num[i] + mod + d // 9+9+0=18 // 9+0+1
		} else {
			sum = mod + d
		}

		d = 0

		if sum > 9 { // 10
			d = 1
			sum %= 10 // 8 // 0
		}

		if i >= 0 {
			num[i] = sum // 8 -> [9,8] // 0 -> [0,8]
		} else {
			num = append([]int{sum}, num...)
		}

		if k > 0 {
			k /= 10 // 0
		}
		i-- // 0 // -1
	}

	if d > 0 {
		if i < 0 {
			num = append([]int{1}, num...) // [1,0,8]
		} else {
			num[i] += d
		}
	}

	return num
}
