package _322_Coin_Change

// [1,2,5]
// 11

// [1,1,2,2,1,2,2,3,3, 2, 3]
//  1 2 3 4 5 6 7 8 9 10 11

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount)
	curr := 0

	for curr < amount {
		for i := 0; i < len(coins); i++ {
			if curr+1 == coins[i] {
				dp[curr] = 1
				break
			}
			if curr+1 < coins[i] || dp[curr-coins[i]] == 0 {
				continue
			}

			if dp[curr] == 0 {
				dp[curr] = dp[curr-coins[i]] + 1
			} else {
				dp[curr] = min(dp[curr], dp[curr-coins[i]]+1)
			}
		}
		curr++
	}

	if dp[amount-1] == 0 {
		return -1
	}

	return dp[amount-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
