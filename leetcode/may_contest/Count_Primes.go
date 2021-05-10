package may_contest

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	counter := 0
	primes := make([]bool, n-2)

	for i := 0; i < n-2; i++ {
		primes[i] = true
	}

	for i := 2; i < n; i++ {
		if primes[i-2] == false {
			continue
		}

		if i*i > n {
			break
		}

		for j := i * i; j < n; j += i {
			primes[j-2] = false
		}
	}

	for i := 0; i < n-2; i++ {
		if primes[i] == true {
			counter++
		}
	}

	return counter
}
