package _645_Set_Mismatch

// [1,2,2,4]
// []

func findErrorNums(nums []int) []int {
	hp := make(map[int]int)
	result := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		hp[nums[i]]++

		if hp[nums[i]] > 1 {
			result[0] = nums[i]
			break
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[abs(nums[i])-1] > 0 {
			nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result[1] = i + 1
			break
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
