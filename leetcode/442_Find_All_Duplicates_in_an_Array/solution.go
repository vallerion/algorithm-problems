package _442_Find_All_Duplicates_in_an_Array

func findDuplicates(nums []int) []int {
	mp := make(map[int]int)
	duplicates := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++

		if mp[nums[i]] > 1 {
			duplicates = append(duplicates, nums[i])
		}
	}

	return duplicates
}
