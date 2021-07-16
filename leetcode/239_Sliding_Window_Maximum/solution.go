package _239_Sliding_Window_Maximum

func maxSlidingWindow(nums []int, k int) []int {
	maxes := make([]int, 0)
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if len(maxes) > 0 && nums[i] > maxes[len(maxes)-1] {
			maxes = []int{}
		}

		for len(maxes) > 0 && nums[i] > maxes[0] {
			maxes = maxes[1:]
		}

		if len(maxes) == 0 || maxes[len(maxes)-1] <= nums[i] {
			maxes = append(maxes, nums[i])
		} else {
			maxes = append([]int{nums[i]}, maxes...)
		}

		if i >= k-1 {
			result = append(result, maxes[len(maxes)-1])

			if maxes[len(maxes)-1] == nums[i-k+1] {
				maxes = maxes[:len(maxes)-1]
			}
		}
	}

	return result
}
