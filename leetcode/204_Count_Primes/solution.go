package _204_Count_Primes

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}

	nums := make([]bool, n)

	for i := 0; i < len(nums); i++ {
		nums[i] = true
	}

	for i := 2; i*i < n; i++ {
		if nums[i] == false {
			continue
		}

		for j := i * i; j < n; j += i {
			nums[j] = false
		}
	}

	count := 0
	for i := 2; i < len(nums); i++ {
		if nums[i] == true {
			count++
		}
	}

	return count
}
