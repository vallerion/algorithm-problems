package _200_Number_of_Islands

func find(set []int, target int) int {
	if set[target] == target {
		return target
	}

	set[target] = find(set, set[target])
	return set[target]
}

func union(set []int, a, b int) {
	ax, bx := find(set, a), find(set, b)

	if ax != bx {
		set[bx] = ax
	}
}

func numIslands(grid [][]byte) int {
	// if len(grid) == 0

	newGrip := make([][]int, 0)
	dsu := make([]int, 0)
	m, n, k := len(grid), len(grid[0]), 0

	for i := 0; i < m; i++ {
		newGrip = append(newGrip, make([]int, n))

		for j := 0; j < n; j++ {
			if (grid[i][j] - 48) == 1 {
				dsu = append(dsu, k)
				newGrip[i][j] = k
				k++
			} else {
				newGrip[i][j] = -1
			}
		}
	}

	//i, j := 0, 0
	//
	//for i < n-1 && j < m-1 {
	//	if newGrip[i][j] < 0 {
	//		continue
	//	}
	//
	//	if newGrip[i][j] > 0 && newGrip[i+1][j] > 0 {
	//		union(dsu, newGrip[i][j], newGrip[i+1][j])
	//	}
	//
	//	if newGrip[i][j] > 0 && newGrip[i][j+1] > 0 {
	//		union(dsu, newGrip[i][j], newGrip[i+1][j])
	//	}
	//
	//	i++
	//}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if newGrip[i][j] < 0 {
				continue
			}

			if i+1 < m && newGrip[i][j] >= 0 && newGrip[i+1][j] >= 0 {
				union(dsu, newGrip[i][j], newGrip[i+1][j])
			}

			if j+1 < n && newGrip[i][j] >= 0 && newGrip[i][j+1] >= 0 {
				union(dsu, newGrip[i][j], newGrip[i][j+1])
			}

			if i-1 >= 0 && newGrip[i][j] >= 0 && newGrip[i-1][j] >= 0 {
				union(dsu, newGrip[i][j], newGrip[i-1][j])
			}

			if j-1 >= 0 && newGrip[i][j] >= 0 && newGrip[i][j-1] >= 0 {
				union(dsu, newGrip[i][j], newGrip[i][j-1])
			}
		}
	}

	dsuCounter := make(map[int]int, 0)

	for i := 0; i < len(dsu); i++ {
		dsuCounter[find(dsu, dsu[i])]++
	}

	//fmt.Println(dsu, dsuCounter)

	return len(dsuCounter)
}
