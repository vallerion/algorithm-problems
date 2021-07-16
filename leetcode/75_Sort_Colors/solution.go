package _75_Sort_Colors

func sortColors(nums []int) {
	hp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hp[nums[i]]++
	}

	for i := 0; i < len(nums); i++ {
		if hp[0] > 0 {
			nums[i] = 0
			hp[0]--
			continue
		}
		if hp[1] > 0 {
			nums[i] = 1
			hp[1]--
			continue
		}
		if hp[2] > 0 {
			nums[i] = 2
			hp[2]--
			continue
		}
	}
}
