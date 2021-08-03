package _122_Best_Time_to_Buy_and_Sell_Stock_II

// [7,1,5,3,6,4]
// 1-7=-6, 0
// 5-1=4, 4

func maxProfit(prices []int) int {
	profit := 0

	for i := 0; i < len(prices)-1; i++ {
		diff := prices[i+1] - prices[i]
		if diff > 0 {
			profit += diff
		}
	}

	return profit
}
