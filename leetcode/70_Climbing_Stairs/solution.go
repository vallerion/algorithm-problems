package _70_Climbing_Stairs

// 0 -> 0
// 1 -> 1 -> 1
// 2 -> 2, 1+1 -> 2
// 3 -> 2+1,1+1+1,1+2 -> 3
// 4 -> 2+2,1+1+1+1,2+1+1,1+1+2,1+2+1 -> 5
// 5 ->

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	first, second := 1, 2

	for i := 2; i < n; i++ {
		second, first = first+second, second
	}

	return second
}
