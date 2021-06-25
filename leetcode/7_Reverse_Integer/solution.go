package _7_Reverse_Integer

import "math"

func reverse(x int) int {
	result := 0
	minus := false

	if x < 0 {
		minus = true
		x = -x
	}

	for x > 0 {
		result *= 10
		d := x % 10

		result += d

		x /= 10
	}

	if minus {
		result = -result
	}

	if result >= math.MaxInt32 || result <= math.MinInt32 {
		result = 0
	}

	return result
}
