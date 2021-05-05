package exercises

// [5,3,43,4,2]
// 3: [5,5,43,4,2] -> [3,5,43,4,2]
// 43: do nothing -> [3,5,43,4,2]
// 4: [3,5,43,43,2] -> [3,5,5,43,2] ->  [3,4,5,43,2]
// 2: [3,4,5,43,43] -> [3,4,5,5,43] -> [3,4,4,5,43] -> [3,3,4,5,43] -> [2,3,4,5,43]

func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		j := i - 1
		for j >= 0 && curr < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = curr
	}

	return nums
}
