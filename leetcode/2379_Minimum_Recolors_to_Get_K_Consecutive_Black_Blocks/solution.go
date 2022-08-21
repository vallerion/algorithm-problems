package _2379_Minimum_Recolors_to_Get_K_Consecutive_Black_Blocks

// blocks = "WBBWWBBWBW", k = 7
// W B B W W B B W B W
// 1 1 1 2 3 3 3 3 3 4
// 0 1 2 3 4 5 6 7 8 9

func minimumRecolors(blocks string, k int) int {
	i := 0
	counter, minCounter := 0, len(blocks)

	for _, s := range blocks {
		if s == 'W' {
			counter++
		}

		if i+1 >= k {
			if minCounter > counter {
				minCounter = counter
			}

			if blocks[i+1-k] == 'W' {
				counter--
			}
		}

		i++
	}

	return minCounter
}
