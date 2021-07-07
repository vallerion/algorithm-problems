package july_contest

import (
	"fmt"
	"sort"
)

func kthSmallest(matrix [][]int, k int) int {
	n, m := len(matrix), len(matrix[0])
	//hm := make(map[int]int)
	nums := make([]int, 0)

	//shift := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//hm[matrix[i][j]]++
			nums = append(nums, matrix[i][j])

			//if hm[matrix[i][j]] == 1 {
			//
			//}

			//if hm[matrix[i][j]] > 1 {
			//	shift++
			//}
		}
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums)

	if k-1 >= len(nums) || k-1 < 0 {
		return -1
	}

	return nums[k-1]
}
