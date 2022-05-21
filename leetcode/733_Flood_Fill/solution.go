package _733_Flood_Fill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	dfs(sr, sc, &image, newColor, image[sr][sc])

	return image
}

func dfs(i, j int, image *[][]int, newColor, oldColor int) {
	if (*image)[i][j] != oldColor {
		return
	}

	(*image)[i][j] = newColor

	if i < len(*image)-1 {
		dfs(i+1, j, image, newColor, oldColor)
	}
	if i > 0 {
		dfs(i-1, j, image, newColor, oldColor)
	}
	if j < len((*image)[0])-1 {
		dfs(i, j+1, image, newColor, oldColor)
	}
	if j > 0 {
		dfs(i, j-1, image, newColor, oldColor)
	}
}
