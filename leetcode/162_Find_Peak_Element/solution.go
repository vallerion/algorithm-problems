package _162_Find_Peak_Element

// [1,2]
//

// [1,2,1,3,5,6,4]

// [1,8,2,3,4,5,6,7]

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	l, r := 0, len(nums)-1
	res := 0

	for l <= r {
		q := (l + r) / 2

		if q-1 >= 0 && nums[q-1] > nums[q] {
			r = q - 1
		} else if q+1 < len(nums) && nums[q+1] > nums[q] {
			l = q + 1
		} else {
			res = q
			break
		}
	}

	return res
}
