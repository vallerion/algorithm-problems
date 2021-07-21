package _378_Kth_Smallest_Element_in_a_Sorted_Matrix

// n * n + m
func kthSmallest(matrix [][]int, k int) int {
	res := make([]int, 0)

	for i := 0; i < len(matrix); i++ {
		res = merge(res, matrix[i])
	}

	return res[k-1]
}

func merge(left, right []int) []int {
	result := make([]int, 0)

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	if i < len(left) {
		result = append(result, left[i:]...)
	}

	if j < len(right) {
		result = append(result, right[j:]...)
	}

	return result
}
