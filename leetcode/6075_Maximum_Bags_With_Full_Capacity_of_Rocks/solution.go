package _6075_Maximum_Bags_With_Full_Capacity_of_Rocks

import (
	"fmt"
	"sort"
)

// capacity = [10,4,2], rocks = [4,2,0], additionalRocks = 6
// [10,2,0] -> 1
// [4,4,2]

// 6 2 2

// [2,3,4,5]
// [1,2,4,4]
// 2
// [10,2,2]
// [2,2,0]
// 100
// [91,54,63,99,24,45,78]
// [35,32,45,98,6,1,25]
// 17

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	for i := 0; i < len(rocks); i++ {
		rocks[i] = capacity[i] - rocks[i]
	}

	fmt.Println(rocks)

	sort.Ints(rocks)
	cnt := 0

	for i := 0; i < len(rocks) && additionalRocks > 0; i++ {
		if additionalRocks < rocks[i] {
			break
		}

		additionalRocks -= rocks[i]
		cnt++
	}

	return cnt
}
