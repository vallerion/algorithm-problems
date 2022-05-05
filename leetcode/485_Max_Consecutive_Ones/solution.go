package _485_Max_Consecutive_Ones

func findMaxConsecutiveOnes(nums []int) int {
	temp, max := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			temp++
		} else {
			temp = 0
		}

		if temp > max {
			max = temp
		}
	}
	return max
}
