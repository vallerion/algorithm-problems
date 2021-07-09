package _53_Maximum_Subarray

// [5,4,-1,7,8]

//  5 or  5+4 -> max=9  -> [5,9]
// -1 or -1+9 -> max=9  -> [5,9,8]
//  7 or  7+8 -> max=15 -> [5,9,8,15]
//  8 or  8+15 -> max=23 -> [5,9,8,15,26]

func maxSubArray(nums []int) int {
	arr := make([]int, len(nums))

	arr[0] = nums[0]
	maxSum := arr[0]

	for i := 1; i < len(nums); i++ {
		arr[i] = max(nums[i], nums[i]+arr[i-1])
		maxSum = max(maxSum, arr[i])
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
