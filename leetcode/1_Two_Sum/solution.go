package _1_Two_Sum

func twoSum(nums []int, target int) []int {
	hp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hp[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if v, ok := hp[target-nums[i]]; ok && v != i {
			return []int{i, v}
		}
	}

	return []int{}
}
