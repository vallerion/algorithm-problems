package _70_Climbing_Stairs

var cache map[int]int

func climbStairs(n int) int {
	if cache == nil {
		cache = make(map[int]int)
	}
	if n <= 2 {
		return n
	}
	if _, ok := cache[n]; ok {
		return cache[n]
	}

	cache[n] = climbStairs(n-1) + climbStairs(n-2)

	return cache[n]
}
