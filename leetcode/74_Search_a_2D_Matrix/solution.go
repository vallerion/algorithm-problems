package _74_Search_a_2D_Matrix

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {

		if binSearch(matrix[i], target) {
			return true
		}
	}

	return false
}

func binSearch(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	} else if len(nums) == 1 {
		return nums[0] == target
	}

	start, end := 0, len(nums)-1

	for start <= end {
		q := start + ((end - start) / 2)

		if nums[q] == target {
			return true
		}

		if nums[q] > target {
			end = q - 1
		} else {
			start = q + 1
		}
	}

	return false
}
