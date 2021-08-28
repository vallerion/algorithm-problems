package _509_Fibonacci_Number

func fib(n int) int {
	if n <= 1 {
		return n
	}

	first, second := 0, 1

	for i := 1; i < n; i++ {
		second, first = first+second, second
	}

	return second
}
