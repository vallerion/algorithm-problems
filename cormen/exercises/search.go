package exercises

// BinarySearch return searched item
func BinarySearch(nums []int, needed int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		if nums[0] == needed {
			return nums[0]
		} else {
			return -1
		}
	}

	q := len(nums) / 2

	if nums[q] == needed {
		return nums[q]
	} else if nums[q] > needed {
		return BinarySearch(nums[:q], needed)
	} else {
		return BinarySearch(nums[q:], needed)
	}
}

// BinarySearchIndex return index of searched item
func BinarySearchIndex(nums []int, start, end, needed int) int {
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
		return BinarySearchIndex(nums, start, q, needed)
	} else {
		return BinarySearchIndex(nums, q+1, end, needed)
	}
}

func LinearSearch(nums []int, needed int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == needed {
			return i
		}
	}

	return -1
}
