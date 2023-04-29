package _48_Rotate_Image

// i+n-1,j -> i,j
// i+n-1,j+n-1 -> i+n-1,j
// i,j+n-1 -> i+n-1,j+n-1
// i,j -> i,j+n-1

func rotate(matrix [][]int) {
	N := len(matrix)

	for i := 0; i < N/2; i++ {
		for j := i; j < N-1-i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[N-1-j][i]
			matrix[N-1-j][i] = matrix[N-1-i][N-1-i]
			matrix[N-1-i][N-1-i] = matrix[j][N-1-i]
			matrix[j][N-1-i] = temp
		}
	}
}
