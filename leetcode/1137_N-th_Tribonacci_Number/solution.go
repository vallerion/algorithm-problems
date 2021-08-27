package _1137_N_th_Tribonacci_Number

func tribonacci(n int) int {
	t1, t2, t3 := 0, 1, 1

	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	}

	for i := 2; i < n; i++ {
		temp1, temp2 := t1, t2
		t1 = t2
		t2 = t3
		t3 = temp1 + temp2
	}

	return t3
}
