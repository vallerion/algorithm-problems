package _26_Remove_Duplicates_from_Sorted_Array

// n = [0,0,1,1,1,2,2,3,3,4], i = 1
// n[i] == n[i-1] -> remove i-1, i -= 1
// ...

func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i:]...)
			i--
		}
	}

	return len(nums)
}