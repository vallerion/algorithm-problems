package _200_Number_of_Islands

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

func (ds *DisjointSet) IsConnected(x, y int) bool {
	return ds.Find(x) == ds.Find(y)
}

func numIslands(grid [][]byte) int {
	R, C := len(grid), len(grid[0])

	ds := Constructor(R * C)

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if i > 0 && grid[i][j] == '1' && grid[i-1][j] == '1' {
				ds.Union(i*len(grid[0])+j, (i-1)*len(grid[0])+j)
			}

			if j > 0 && grid[i][j] == '1' && grid[i][j-1] == '1' {
				ds.Union(i*len(grid[0])+j, i*len(grid[0])+(j-1))
			}
		}
	}

	islands := make(map[int]int)

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == '1' {
				islands[ds.Find(i*len(grid[0])+j)]++
			}
		}
	}

	return len(islands)
}
