package april_challenge

// 0 1 1 2 3 5 8 13 21
// 0 1 2 3 4 5 6  7  8
func fib(n int) int {
	if n == 0 {
		return 0
	}

	first := 0
	second := 1
	i := 1

	for i < n {
		first, second = second, first + second
		i++
	}

	return second
}
