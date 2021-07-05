package july_contest

func matrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])

	if r*c != n*m {
		return mat
	}

	newMat := make([][]int, r)

	x, y := 0, 0

	for i := 0; i < r; i++ {
		newMat[i] = make([]int, c)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			newMat[x][y] = mat[i][j]
			y++

			if y == c {
				y = 0
				x++
			}
		}
	}

	return newMat
}
