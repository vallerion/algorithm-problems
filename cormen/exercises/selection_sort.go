package exercises

// [4,3,1,5]
// [1,3,4,5]

func SelectionSort(nums []int) []int {
	currMinIndex := 0

	for i := 0; i < len(nums); i++ {
		currMinIndex = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[currMinIndex] && j != currMinIndex {
				currMinIndex = j
			}
		}
		nums[i], nums[currMinIndex] = nums[currMinIndex], nums[i]
	}

	return nums
}
