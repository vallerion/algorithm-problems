package _215_Kth_Largest_Element_in_an_Array

import "sort"

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)

	if k > len(nums) {
		return -1
	}

	return nums[len(nums)-k]
}
