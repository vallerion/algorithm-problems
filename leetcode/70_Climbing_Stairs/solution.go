package _70_Climbing_Stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	steps := make([]int, n)

	steps[0] = 1
	steps[1] = 2

	for i := 2; i < n; i++ {
		steps[i] = steps[i-1] + steps[i-2]
	}

	return steps[n-1]
}

func climbStairs_Recursive(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	return climbStairs(n-2) + climbStairs(n-1)
}
