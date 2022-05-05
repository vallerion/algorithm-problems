package _746_Min_Cost_Climbing_Stairs

// [100,1,1,1,1,100,1,1,100,1]
// [100,1,2,2,3,102,4,5,104,6]

// [1,100,1,1,1,100,1,1,100,1]
// [1,1  ,2,3,3,4  ,4,5,5  ,6]

// [10,15,20]
// [10,15,20]

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}

	hh := make([]int, len(cost))
	hh[0] = cost[0]
	hh[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		hh[i] = min(hh[i-1], hh[i-2]) + cost[i]

		if i == len(cost)-1 {
			hh[i] = min(hh[i], hh[i-1])
		}
	}

	return hh[len(hh)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
