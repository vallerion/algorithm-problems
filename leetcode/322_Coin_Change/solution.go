package _322_Coin_Change

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	st := make([]int, 0)
	st = append(st, amount)
	minLevels := amount + 1

	for len(st) > 0 {
		n := len(st)

		levels := 0
		for i := 0; i < n; i++ {
			for j := 0; j < len(coins); j++ {
				div := st[i] / coins[j]
				levels += div
				rest := st[i] % coins[j]

				if rest == 0 && minLevels > levels {
					minLevels = levels
					continue
				}
				if rest < 0 {
					continue
				}

				st = append(st, rest)
			}
		}

		st = st[n:]
	}

	if minLevels < amount+1 {
		return minLevels
	}

	return -1
}
