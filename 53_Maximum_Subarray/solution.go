package _53_Maximum_Subarray

// [-2,1,-3,4,-1,2,1,-5,4]
// curr=curr-2, curr=curr<0?0:curr -> curr=2, curr>max?max=curr // -2
// curr=1, max=1 // 1
// curr=-2 -> curr=0, max=1 // -3
// curr=4, max=4 // 4
// curr=3, max=4 // -1
// curr=5, max=5 // 2
// curr=6, max=6 // 1
// curr=1, max=6 // -5

// [-2,-1]
// curr=-2, max=-2, curr=0
// curr=-1, max=-1, curr=0
func maxSubArray(nums []int) int {
	curr := 0
	max := 0

	if len(nums) > 0 {
		curr = nums[0]
		max = curr
	}

	for i := 0; i < len(nums); i++ {
		curr = curr + nums[i]
		if curr > max {
			max = curr
		}
		if curr < 0 {
			curr = 0
		}
	}

	return max
}

// okay, lets brute force
// [-2,-1]
// max=-2, tempSum=-3
//
// [-1,0,-2]
// max=-1, tempSum=-1 -> tempSum=-3
// tempSum=0, max=0
func maxSubArrayBF(nums []int) int {
	max := 0

	if len(nums) > 0 {
		max = nums[0]
	}

	for i := 0; i < len(nums); i++ {
		tempSum := nums[i]
		if tempSum > max {
			max = tempSum
		}
		for j := i + 1; j < len(nums); j++ {
			tempSum = tempSum + nums[j]
			if tempSum > max {
				max = tempSum
			}
		}
	}

	return max
}
