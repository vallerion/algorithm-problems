package _198_House_Robber

// [2,1,1,2,1,2,2,2]
// [8,7,5,6,3,4,2,2]

func rob(nums []int) int {
	maxAmount := 0

	for i := len(nums) - 1; i >= 0; i-- {
		if i+3 < len(nums) {
			nums[i] = nums[i] + max(nums[i+2], nums[i+3])
		} else if i+2 < len(nums) {
			nums[i] += nums[i+2]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > maxAmount {
			maxAmount = nums[i]
		}
	}

	return maxAmount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
