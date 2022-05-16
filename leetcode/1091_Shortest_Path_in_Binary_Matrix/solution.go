package _1091_Shortest_Path_in_Binary_Matrix

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	visited := make([][]int, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, len(grid))
	}
	visited[0][0] = 1

	st := make([][]int, 0)
	st = append(st, []int{0, 0})

	for len(st) > 0 {
		n := len(st)

		for i := 0; i < n; i++ {
			// upper row
			if st[i][0] > 0 && st[i][1] > 0 && grid[st[i][0]-1][st[i][1]-1] == 0 {
				if visited[st[i][0]-1][st[i][1]-1] == 0 || (visited[st[i][0]-1][st[i][1]-1] > 0 && visited[st[i][0]][st[i][1]]+1 < visited[st[i][0]-1][st[i][1]-1]) {
					st = append(st, []int{st[i][0] - 1, st[i][1] - 1})
					visited[st[i][0]-1][st[i][1]-1] = visited[st[i][0]][st[i][1]] + 1
				}
			}
			if st[i][0] > 0 && grid[st[i][0]-1][st[i][1]] == 0 {
				if visited[st[i][0]-1][st[i][1]] == 0 || (visited[st[i][0]-1][st[i][1]] > 0 && visited[st[i][0]][st[i][1]]+1 < visited[st[i][0]-1][st[i][1]]) {
					st = append(st, []int{st[i][0] - 1, st[i][1]})
					visited[st[i][0]-1][st[i][1]] = visited[st[i][0]][st[i][1]] + 1
				}
			}
			if st[i][0] > 0 && st[i][1] < len(grid)-1 && grid[st[i][0]-1][st[i][1]+1] == 0 {
				if visited[st[i][0]-1][st[i][1]+1] == 0 || (visited[st[i][0]-1][st[i][1]+1] > 0 && visited[st[i][0]][st[i][1]]+1 < visited[st[i][0]-1][st[i][1]+1]) {
					st = append(st, []int{st[i][0] - 1, st[i][1] + 1})
					visited[st[i][0]-1][st[i][1]+1] = visited[st[i][0]][st[i][1]] + 1
				}
			}

			// bottom row
			if st[i][0] < len(grid)-1 && st[i][1] > 0 && grid[st[i][0]+1][st[i][1]-1] == 0 {
				if visited[st[i][0]+1][st[i][1]-1] == 0 || (visited[st[i][0]+1][st[i][1]-1] > 0 && visited[st[i][0]][st[i][1]]+1 < visited[st[i][0]+1][st[i][1]-1]) {
					st = append(st, []int{st[i][0] + 1, st[i][1] - 1})
					visited[st[i][0]+1][st[i][1]-1] = visited[st[i][0]][st[i][1]] + 1
				}
			}
			if st[i][0] < len(grid)-1 && grid[st[i][0]+1][st[i][1]] == 0 {
				if visited[st[i][0]+1][st[i][1]] == 0 || (visited[st[i][0]+1][st[i][1]] > 0 && visited[st[i][0]][st[i][1]]+1 < visited[st[i][0]+1][st[i][1]]) {
					st = append(st, []int{st[i][0] + 1, st[i][1]})
					visited[st[i][0]+1][st[i][1]] = visited[st[i][0]][st[i][1]] + 1
				}
			}
			if st[i][0] < len(grid)-1 && st[i][1] < len(grid)-1 && grid[st[i][0]+1][st[i][1]+1] == 0 {
				if visited[st[i][0]+1][st[i][1]+1] == 0 || (visited[st[i][0]+1][st[i][1]+1] > 0 && visited[st[i][0]][st[i][1]]+1 < visited[st[i][0]+1][st[i][1]+1]) {
					st = append(st, []int{st[i][0] + 1, st[i][1] + 1})
					visited[st[i][0]+1][st[i][1]+1] = visited[st[i][0]][st[i][1]] + 1
				}
			}

			// left
			if st[i][1] > 0 && grid[st[i][0]][st[i][1]-1] == 0 {
				if visited[st[i][0]][st[i][1]-1] == 0 || (visited[st[i][0]][st[i][1]-1] > 0 && visited[st[i][0]][st[i][1]]+1 < visited[st[i][0]][st[i][1]-1]) {
					st = append(st, []int{st[i][0], st[i][1] - 1})
					visited[st[i][0]][st[i][1]-1] = visited[st[i][0]][st[i][1]] + 1
				}
			}

			// right
			if st[i][1] < len(grid)-1 && grid[st[i][0]][st[i][1]+1] == 0 {
				if visited[st[i][0]][st[i][1]+1] == 0 || (visited[st[i][0]][st[i][1]+1] > 0 && visited[st[i][0]][st[i][1]]+1 < visited[st[i][0]][st[i][1]+1]) {
					st = append(st, []int{st[i][0], st[i][1] + 1})
					visited[st[i][0]][st[i][1]+1] = visited[st[i][0]][st[i][1]] + 1
				}
			}

		}

		st = st[n:]
	}

	if visited[len(visited)-1][len(visited)-1] == 0 {
		return -1
	}

	return visited[len(visited)-1][len(visited)-1]
}
