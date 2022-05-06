package _905_Sort_Array_By_Parity

// [8,3,1,2,4,4]
// [8,4,4,2,1,3]

// even -> left, odd -> right
func sortArrayByParity(nums []int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		if nums[l]%2 != 0 && nums[r]%2 == 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		} else if nums[l]%2 == 0 {
			l++
		} else {
			r--
		}
	}

	return nums
}
