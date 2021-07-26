package _283_Move_Zeroes

// [0,1,0,3,12]
// [1,0,0,3,12]
// [1,3,12,0,0]

// [1,0,1]
// [1,1,0]
func moveZeroes(nums []int) {
	i, j := 0, 1

	for j < len(nums) {
		if nums[i] == 0 && nums[j] != 0 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			i++
			j++
		} else if nums[i] == 0 && nums[j] == 0 {
			j++
		} else {
			i++
			j++
		}
	}
}
