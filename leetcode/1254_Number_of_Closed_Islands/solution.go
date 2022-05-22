package _1254_Number_of_Closed_Islands

type DisjointSet struct {
	values, rank []int
}

func Constructor(length int) *DisjointSet {
	values, rank := make([]int, length), make([]int, length)
	for i := 0; i < length; i++ {
		values[i] = i
	}

	return &DisjointSet{values, rank}
}

func (ds *DisjointSet) Find(x int) int {
	if ds.values[x] == x {
		return x
	}

	ds.values[x] = ds.Find(ds.values[x])
	return ds.values[x]
}

func (ds *DisjointSet) Union(x, y int) {
	rootX, rootY := ds.Find(x), ds.Find(y)

	if rootX != rootY {
		if ds.rank[rootY] > ds.rank[rootX] {
			ds.values[rootX] = rootY
		} else if ds.rank[rootY] < ds.rank[rootX] {
			ds.values[rootY] = rootX
		} else {
			ds.values[rootY] = rootX
			ds.rank[rootX]++
		}
	}
}

func closedIsland(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	ds := Constructor(R * C)

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == 0 && i > 0 && grid[i-1][j] == 0 {
				ds.Union(i*C+j, (i-1)*C+j)
			}
			if grid[i][j] == 0 && j > 0 && grid[i][j-1] == 0 {
				ds.Union(i*C+j, i*C+j-1)
			}
		}
	}

	exclude := make(map[int]int)

	for j := 0; j < C; j++ {
		if grid[0][j] == 0 {
			exclude[ds.Find(j)]++
		}
	}

	for j := 0; j < C; j++ {
		if grid[R-1][j] == 0 {
			exclude[ds.Find((R-1)*C+j)]++
		}
	}

	for i := 0; i < R; i++ {
		if grid[i][0] == 0 {
			exclude[ds.Find(i*C)]++
		}
	}

	for i := 0; i < R; i++ {
		if grid[i][C-1] == 0 {
			exclude[ds.Find(i*C+C-1)]++
		}
	}

	islands := make(map[int]int)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == 1 {
				continue
			}
			_, ok := exclude[ds.Find(i*C+j)]
			if !ok {
				islands[ds.Find(i*C+j)]++
			}
		}
	}

	return len(islands)
}
