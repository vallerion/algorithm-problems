package _329_Longest_Increasing_Path_in_a_Matrix

func longestIncreasingPath(matrix [][]int) int {
	longest := make([][]int, len(matrix))
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(longest); i++ {
		longest[i] = make([]int, len(matrix[0]))
		visited[i] = make([]bool, len(matrix[0]))
	}

	dfs(0, 0, 0, 0, matrix, &visited, &longest)

	max := 0

	for i := 0; i < len(longest); i++ {
		for j := 0; j < len(longest[i]); j++ {
			if longest[i][j] > max {
				max = longest[i][j]
			}
		}
	}

	return max + 1
}

func dfs(prevI, prevJ, i, j int, matrix [][]int, visited *[][]bool, longest *[][]int) {
	(*visited)[i][j] = true

	if matrix[i][j] > matrix[prevI][prevJ] && (*longest)[prevI][prevJ]+1 > (*longest)[i][j] {
		(*longest)[i][j] = (*longest)[prevI][prevJ] + 1
	}

	if i > 0 && (!(*visited)[i-1][j] || ((*visited)[i-1][j] && matrix[i-1][j] > matrix[i][j] && (*longest)[i][j]+1 > (*longest)[i-1][j])) {
		dfs(i, j, i-1, j, matrix, visited, longest)
	}
	if j < len(matrix[0])-1 && (!(*visited)[i][j+1] || ((*visited)[i][j+1] && matrix[i][j+1] > matrix[i][j] && (*longest)[i][j]+1 > (*longest)[i][j+1])) {
		dfs(i, j, i, j+1, matrix, visited, longest)
	}
	if i < len(matrix)-1 && (!(*visited)[i+1][j] || ((*visited)[i+1][j] && matrix[i+1][j] > matrix[i][j] && (*longest)[i][j]+1 > (*longest)[i+1][j])) {
		dfs(i, j, i+1, j, matrix, visited, longest)
	}
	if j > 0 && (!(*visited)[i][j-1] || ((*visited)[i][j-1] && matrix[i][j-1] > matrix[i][j] && (*longest)[i][j]+1 > (*longest)[i][j-1])) {
		dfs(i, j, i, j-1, matrix, visited, longest)
	}
}
