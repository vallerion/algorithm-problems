package _73_Set_Matrix_Zeroes

func setZeroes(matrix [][]int) {
	cols, rows := make(map[int]int), make(map[int]int)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows[i]++
				cols[j]++
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			_, rowZero := rows[i]
			_, colZero := cols[j]
			if rowZero == true {
				matrix[i][j] = 0
			}
			if rowZero == false && colZero == true {
				matrix[i][j] = 0
			}
		}
	}
}
