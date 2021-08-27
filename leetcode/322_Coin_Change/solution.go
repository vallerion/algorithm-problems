package _322_Coin_Change

import "math"

// coins = [1,2,5], amount = 11
// [0,0,0,0,0,0,0,0,0,0,0]
// [0,1,0,0,0,0,0,0,0,0,0]
// [0,1,1,2,2,0,0,0,0,0,0]

// 1 -> 1
// 2 -> 1
// 3 -> 2
// 4 -> 2

func coinChange(coins []int, amount int) int {
	changes := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		smallest := math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if coins[j] < i {
				smallest = min(smallest, changes[i-coins[j]]+1)
			} else if coins[j] == i {
				smallest = 1
			}
		}

		changes[i] = smallest
	}

	if changes[amount] == math.MaxInt32 {
		return -1
	}

	return changes[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
