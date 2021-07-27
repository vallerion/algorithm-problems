package _53_Maximum_Subarray

func maxSubArray(nums []int) int {
	sums := make([]int, len(nums))

	maxSum := nums[0]

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sums[i] = nums[i]
			continue
		}

		sums[i] = max(sums[i-1]+nums[i], nums[i])
		if sums[i] > maxSum {
			maxSum = sums[i]
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
