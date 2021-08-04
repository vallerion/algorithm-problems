package _62_Unique_Paths

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	matrix[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			top := 0
			left := 0

			if i > 0 {
				top = matrix[i-1][j]
			}
			if j > 0 {
				left = matrix[i][j-1]
			}
			matrix[i][j] = top + left
		}
	}

	return matrix[m-1][n-1]
}
