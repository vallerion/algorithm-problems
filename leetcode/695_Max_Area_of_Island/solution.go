package _695_Max_Area_of_Island

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

func maxAreaOfIsland(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	ds := Constructor(R * C)
	atLeastOneIsland := false

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == 1 {
				if i > 0 && grid[i-1][j] == 1 {
					ds.Union(i*C+j, (i-1)*C+j)
				}
				if j > 0 && grid[i][j-1] == 1 {
					ds.Union(i*C+j, i*C+j-1)
				}

				atLeastOneIsland = true
			}
		}
	}

	if !atLeastOneIsland {
		return 0
	}

	islands := make(map[int]int)
	max := 0

	for i := 0; i < len(ds.values); i++ {
		root := ds.Find(i)
		islands[root]++

		if islands[root] > max {
			max = islands[root]
		}
	}

	return max
}
