package _55_Jump_Game

// [0,1,2,3,4]
// [2,0]

// -- li=1
//

func canJump(nums []int) bool {

	li := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= li-i {
			li = i
		}
	}

	return li == 0
}
