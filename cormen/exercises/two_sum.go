package exercises

// Given array of integers and value.
// How find pair of items from array,
// sum of their would be equal given value.
// Solve it in O(n log n).
func twoSum(nums []int, target int) bool {
	nums = mergeSort(nums) // O(n log n)

	// Binary search inside for give O(n log n), looks like solved
	// Totally worse time might be O(n log n + n log n).
	for i := 0; i < len(nums) && nums[i] < target; i++ { // O (n)
		j := binarySearch(nums, i+1, len(nums)-1, target-nums[i]) // O(log n)
		if j != -1 {
			return true
		}
	}

	return false
}

func merge(first, second []int) []int {
	result := make([]int, len(first)+len(second))
	i := 0
	j := 0

	for {
		if i == len(first) && j == len(second) {
			break
		}

		if i == len(first) {
			result[i+j] = second[j]
			j++
			continue
		}

		if j == len(second) {
			result[i+j] = first[i]
			i++
			continue
		}

		if first[i] < second[j] {
			result[i+j] = first[i]
			i++
			continue
		} else {
			result[i+j] = second[j]
			j++
			continue
		}
	}

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	q := len(arr) / 2

	left := mergeSort(arr[:q])
	right := mergeSort(arr[q:])

	return merge(left, right)
}

func binarySearch(nums []int, start, end, needed int) int {
	if start >= end {
		if nums[start] == needed {
			return start
		} else {
			return -1
		}
	}

	q := start + ((end - start) / 2)

	if nums[q] == needed {
		return q
	} else if nums[q] > needed {
		return binarySearch(nums, start, q, needed)
	} else {
		return binarySearch(nums, q+1, end, needed)
	}
}
