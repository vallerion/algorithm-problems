package _26_Remove_Duplicates_from_Sorted_Array

// [0,0,0,0,0,1,1,1,2,2,3,3,4]
// [0,1,2,0,0,1,1,1,2,2,3,3,4]

// 0 1
// 0 2
// 0 3
// 0 4
// 0 5 nums[i+1]=nums[j]
// 1 6
// 1 7
// 1 8 nums[i+1]=nums[j]

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	i, j := 0, 1

	for j < len(nums) {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}

	return i + 1
}
