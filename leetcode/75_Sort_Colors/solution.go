package _75_Sort_Colors

func sortColors(nums []int) {
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}

	for i := 0; i < len(nums); i++ {
		if mp[0] > 0 {
			nums[i] = 0
			mp[0]--
		} else if mp[1] > 0 {
			nums[i] = 1
			mp[1]--
		} else {
			nums[i] = 2
			mp[2]--
		}
	}
}
