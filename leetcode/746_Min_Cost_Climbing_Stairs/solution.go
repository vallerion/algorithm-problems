package _746_Min_Cost_Climbing_Stairs

// [1,2,3]
// [3,2,3]

// [100,1,1,2,3,1,100]
// [  0,104,104,103,103,101,0]
// [100,1,1,2,103,101,100]
// 1 + 2 + 1 + 100 = 104

//

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	if len(cost) == 1 {
		return cost[0]
	}

	trace := make([]int, len(cost))

	for i := len(cost) - 1; i >= 0; i-- {
		if i >= len(cost)-2 {
			trace[i] = cost[i]
		} else {
			trace[i] = cost[i] + min(trace[i+1], trace[i+2])
		}
	}

	return min(trace[0], trace[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
