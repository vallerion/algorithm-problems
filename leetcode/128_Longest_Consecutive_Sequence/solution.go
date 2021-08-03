package _128_Longest_Consecutive_Sequence

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	hp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hp[nums[i]]++
	}

	maxLength := 1

	for i := 0; i < len(nums); i++ {
		if _, ok := hp[nums[i]+1]; ok {
			continue
		}
		length := 1

		j := nums[i] - 1
		_, ok := hp[j]
		for ok {
			j--
			length++
			_, ok = hp[j]
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
