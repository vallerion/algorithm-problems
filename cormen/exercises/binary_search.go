package exercises

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

func Bsi(nums []int, start, end, needed int) int {
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
		return Bsi(nums, start, q, needed)
	} else {
		return Bsi(nums, q+1, end, needed)
	}
}
