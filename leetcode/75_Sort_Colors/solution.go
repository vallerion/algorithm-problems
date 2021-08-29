package _75_Sort_Colors

func sortColors(nums []int) {
	zero, one, two := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero++
		}
		if nums[i] == 1 {
			one++
		}
		if nums[i] == 2 {
			two++
		}
	}

	i := 0
	for zero > 0 {
		nums[i] = 0
		i++
		zero--
	}
	for one > 0 {
		nums[i] = 1
		i++
		one--
	}
	for two > 0 {
		nums[i] = 2
		i++
		two--
	}
}
