package _5299_Find_the_K_Beauty_of_a_Number

import "strconv"

// 240
// 2
// 430043
// 2
// 1
// 1
// 22
// 1

func divisorSubstrings(num int, k int) int {
	str := strconv.Itoa(num)

	if len(str) == k {
		return 1
	}

	res := 0
	for i := 0; i < len(str)-k+1; i++ {
		d, _ := strconv.Atoi(str[i : i+k])
		if d == 0 {
			continue
		}

		if num%d == 0 {
			res++
		}
	}

	return res
}

// [10,4,-8,7]
// i=0
