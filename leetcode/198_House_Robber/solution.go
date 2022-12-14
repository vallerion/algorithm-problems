package _198_House_Robber

// [13,11,12,4,9,3,3,1]
// [12,12, 9,9,3,3,1,1]

// [2,1,1,2]
// [2,2,2,0]

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	accum := make([]int, len(nums))

	maxValue := 0
	for i := len(nums) - 1; i >= 0; i-- {
		accum[i] = maxValue
		if i == len(nums)-1 {
			accum[i] = nums[i]
			maxValue = max(maxValue, nums[i])
			continue
		}
		if i < len(nums)-2 {
			nums[i] += max(nums[i+2], accum[i+2])
		}
		maxValue = max(maxValue, nums[i])
	}

	return max(nums[0], nums[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
