package _70_Climbing_Stairs

// 1 -> 1
// 2 -> 1+1 2
// 3 -> 1+1+1 2+1 1+2
// 4 -> 2+2 1+1+1+1 2+1+1 1+2+1 1+1+2

func climbStairs(n int) int {
	prev1, prev2 := 1, 2

	if n <= 3 {
		return n
	}

	for i := 2; i < n; i++ {
		prev2, prev1 = prev1+prev2, prev2
	}

	return prev2
}
