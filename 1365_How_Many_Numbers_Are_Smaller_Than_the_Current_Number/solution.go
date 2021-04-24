package _1365_How_Many_Numbers_Are_Smaller_Than_the_Current_Number

// [8,1,2,2,3]
// 8, 1

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				res[i]++
			}
		}
	}

	return res
}