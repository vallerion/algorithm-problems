package _455_Assign_Cookies

import (
	"sort"
)

// [1,2,3]
// [1,1,3]

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j := 0, 0
	fed := 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			fed++
		}
		j++
	}

	return fed
}
