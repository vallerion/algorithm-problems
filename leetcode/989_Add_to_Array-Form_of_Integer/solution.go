package _989_Add_to_Array_Form_of_Integer

import "strconv"

// num = [1, 2, 3, 4], k = 1
// [1, 2, 3, 4]		[1]
// 4+1 = 5 -> [1, 2, 3, 5]
// 3+0 = 3 -> [1, 2, 3, 5]
// ...

// num = [1, 2, 3, 4], k = 1449
// [1, 2, 3, 4]		[1, 4, 4, 9]
// 4+9+0 = 13 -> 1 / [1, 2, 3, 3]
// 3+4+1 = 8 -> 0 / [1, 2, 8, 3]
// ...

func AddToArrayForm(num []int, k int) []int {
	kk := []rune(strconv.Itoa(k))
	result := make([]int, 0)

	i := len(num) - 1
	j := len(kk) - 1

	sum := 0
	div := 0
	mod := 0

	for i >= 0 || j >= 0 {
		d1 := 0
		d2 := 0

		if i >= 0 {
			d1 = num[i]
		}
		if j >= 0 {
			d2 = int(kk[j] - '0')
		}

		sum = d1 + d2 + div
		mod = sum % 10
		div = sum / 10

		if sum > 9 {
			result = append([]int{mod}, result...)
		} else {
			result = append([]int{sum}, result...)
		}

		i--
		j--

		if i < 0 && j < 0 && div != 0 {
			result = append([]int{div}, result...)
		}
	}

	return result
}
