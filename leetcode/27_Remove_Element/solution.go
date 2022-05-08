package _27_Remove_Element

// [2,2,2,3]

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	cnt := 0

	for i < j {
		if nums[j] == val {
			j--
			cnt++
			continue
		}

		if nums[i] == val {
			nums[i] = nums[j]
			i++
			j--
			cnt++
		} else {
			i++
		}

	}

	return len(nums) - cnt
}
