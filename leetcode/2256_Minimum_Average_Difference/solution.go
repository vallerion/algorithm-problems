package _2256_Minimum_Average_Difference

import "math"

// [2,5,3,9,5,3]
// [2,3,3,4,4,4]
// [5,5,5,4,3,0]
//

func minimumAverageDifference(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	l2r, r2l := make([]int, len(nums)), make([]int, len(nums))
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		l2r[i] = sum / (i + 1)
	}

	sum = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		r2l[i] = sum / (len(nums) - i - 1)
		sum += nums[i]
	}

	min, minIndex := math.MaxInt, 0

	for i := 0; i < len(nums); i++ {
		diff := abs(l2r[i] - r2l[i])
		if min > diff {
			min = diff
			minIndex = i
		}
	}

	return minIndex
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
