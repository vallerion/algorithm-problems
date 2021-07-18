package _7_Reverse_Integer

import "math"

// 123 -> 321

// x=123 -> res=0,  d=3, x=12, res=3
// x=12  -> res=30, d=2, x=1,  res=32
// x=1   -> res=320,d=1, x=0,  res=321

func reverse(x int) int {
	res := 0
	minus := false

	if x < 0 {
		minus = true
		x = -x
	}

	for x > 0 {
		res *= 10

		d := x % 10
		x /= 10

		res += d
	}

	if res > math.MaxInt32 {
		return 0
	}

	if minus == true {
		return -res
	}

	return res
}
