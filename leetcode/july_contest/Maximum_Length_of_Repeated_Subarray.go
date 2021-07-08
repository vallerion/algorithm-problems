package july_contest

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}

	maxLength := 0

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1

				maxLength = max(maxLength, dp[i+1][j+1])
			}
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
