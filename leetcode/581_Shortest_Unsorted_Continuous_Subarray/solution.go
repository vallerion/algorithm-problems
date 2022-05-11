package _581_Shortest_Unsorted_Continuous_Subarray

import "sort"

// 1,3,2,4
// 1,2,3,4
// 0 1 2

// 2,6,4,8,10,9,15
// 2,4,6,8,9,10,15
// 0 1 2 3 4 5 6

// O(n log n)
func findUnsortedSubarrayOnlogn(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	first, second := -1, -1

	for i := 0; i < len(nums); i++ {
		if nums[i] != sorted[i] {
			if first < 0 {
				first = i
			} else {
				second = i
			}
		}
	}

	if first == -1 || second == -1 {
		return 0
	}

	return second - first + 1
}
