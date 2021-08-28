package _1137_N_th_Tribonacci_Number

func tribonacci(n int) int {
	first, second, third := 0, 1, 1

	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	}

	for i := 2; i < n; i++ {
		third, first, second = first+second+third, second, third
	}

	return third
}
