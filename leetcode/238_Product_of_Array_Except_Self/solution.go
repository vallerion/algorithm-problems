package _238_Product_of_Array_Except_Self

// [1,2,3,4]
// leftToRight=[1,1,2,6]
// rightToLeft=[24,12,4,1]
// result=[24,12,8,6]

// [-1,1,0,-3,3]
// leftToRight=[1,-1,-1,0,0]
// rightToLeft=[0,0,-9,3,1]
// result=[0,0,-9,0,0]

func productExceptSelf(nums []int) []int {
	leftToRight, rightToLeft := make([]int, len(nums)), make([]int, len(nums))

	leftToRight[0] = 1
	rightToLeft[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		leftToRight[i] = leftToRight[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		rightToLeft[i] = rightToLeft[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		leftToRight[i] *= rightToLeft[i]
	}

	return leftToRight
}
