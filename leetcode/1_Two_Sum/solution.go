package _1_Two_Sum

// nums = [2,7,11,15], target = 9
// i=0, j=0, 2+7=9, if 9 == 9 return [0,0]

// nums = [3,2,4], target = 6
// i=0, j=1, 3+2=5, -
// i=0, j=2, 3+4=7, -
// i=1, j=2, 2+4=6, +

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
