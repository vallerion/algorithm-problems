package _74_Search_a_2D_Matrix

func searchMatrix(matrix [][]int, target int) bool {

	i, j := 0, len(matrix[0])-1

	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if target < matrix[i][j] {
			j--
		} else {
			i++
		}
	}

	return false
}
