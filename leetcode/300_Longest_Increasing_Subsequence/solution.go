package _300_Longest_Increasing_Subsequence

// [10,9,2,5,3,7,101,18]
// [ 1,1,1,2,2,3,  4, 4]

// [0,1,0,3,2,3]
// [1,2,1,3,3,4]

func lengthOfLIS(nums []int) int {
	ln := make([]int, len(nums))
	maxLen := 1

	for i := 0; i < len(nums); i++ {
		ln[i] = 1
		tmax := 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				tmax = max(tmax, ln[i]+ln[j])
			}
		}

		ln[i] = tmax

		if ln[i] > maxLen {
			maxLen = ln[i]
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
