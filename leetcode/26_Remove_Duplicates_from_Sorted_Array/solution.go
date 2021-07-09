package _26_Remove_Duplicates_from_Sorted_Array

// [0,0,1,1,1,2,2,3,3,4]
// j = 1, i = 0 -> nums[j-1] != nums[i] { nums[j] = nums[i] j++ } -> 0 != 0 -> false
// j = 1, i = 1 -> nums[j-1] != nums[i] { nums[j] = nums[i] j++ } -> 0 != 0 -> false
// j = 1, i = 2 -> nums[j-1] != nums[i] { nums[j] = nums[i] j++ } -> 0 != 1 -> true
// [0,1,1,1,1,2,2,3,3,4]
// j = 2, i = 3 -> nums[j-1] != nums[i] { nums[j] = nums[i] } -> 1 != 1 -> false
// j = 2, i = 4 -> nums[j-1] != nums[i] { nums[j] = nums[i] } -> 1 != 1 -> false
// j = 2, i = 5 -> nums[j-1] != nums[i] { nums[j] = nums[i] } -> 1 != 2 -> true
// [0,1,2,1,1,2,2,3,3,4]
// j = 3, i = 6 -> nums[j-1] != nums[i] { nums[j] = nums[i] } -> 2 != 2 -> false
// j = 3, i = 7 -> nums[j-1] != nums[i] { nums[j] = nums[i] } -> 2 != 3 -> true
// [0,1,2,3,1,2,2,3,3,4]
// j = 4, i = 8 -> nums[j-1] != nums[i] { nums[j] = nums[i] } -> 3 != 3 -> false
// j = 4, i = 9 -> nums[j-1] != nums[i] { nums[j] = nums[i] } -> 3 != 4 -> true
// [0,1,2,3,4,2,2,3,3,4]

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	j := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
