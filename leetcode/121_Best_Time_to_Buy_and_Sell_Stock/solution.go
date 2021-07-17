package _121_Best_Time_to_Buy_and_Sell_Stock

// [7,1,5,3,6,4]

func maxProfit(prices []int) int {
	profit := 0

	i, j := 0, 1

	for j < len(prices) {
		diff := prices[j] - prices[i]

		if diff <= 0 {
			i = j
			j++
			continue
		}

		if diff > profit {
			profit = diff
		}
		j++
	}

	return profit
}
