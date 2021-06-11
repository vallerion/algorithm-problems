package _128_Longest_Consecutive_Sequence

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mp := make(map[int]bool, len(nums))
	largestStreak, currentStreak := 1, 1

	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = true
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]+1]; ok {
			continue
		}

		currentStreak = 1
		currentNum := nums[i] - 1

		for {
			_, ok := mp[currentNum]
			if ok == false {
				break
			}

			currentStreak++
			currentNum--
		}

		largestStreak = max(largestStreak, currentStreak)
	}

	return largestStreak
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
