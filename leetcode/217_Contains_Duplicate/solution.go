package _217_Contains_Duplicate


func containsDuplicate(nums []int) bool {
	hp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hp[nums[i]]++

		if hp[nums[i]] > 1 {
			return true
		}
	}

	return false
}