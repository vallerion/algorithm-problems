package _2270_Number_of_Ways_to_Split_Array

// [10,4,-8,7]
// [2,3,1,0]
// [-100,34,5,-1,33]
func waysToSplitArray(nums []int) int {
	r2l, l2r := make([]int, len(nums)), make([]int, len(nums))
	r2l[0] = nums[0]
	l2r[len(nums)-1] = nums[len(nums)-1]

	for i := 1; i < len(nums)-1; i++ {
		r2l[i] = r2l[i-1] + nums[i]
		l2r[len(nums)-i-1] = l2r[len(nums)-i] + nums[len(nums)-i-1]
	}

	res := 0
	for i := 0; i < len(nums)-1; i++ {
		if r2l[i] >= l2r[i+1] {
			res++
		}
	}

	return res
}
