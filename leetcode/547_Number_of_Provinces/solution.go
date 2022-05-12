package _547_Number_of_Provinces

type DisjointSet struct {
	values []int
}

func Constructor(length int) *DisjointSet {
	values := make([]int, length)
	for i := 0; i < length; i++ {
		values[i] = i
	}

	return &DisjointSet{values}
}

func (ds *DisjointSet) Find(x int) int {
	if ds.values[x] == x {
		return x
	}

	ds.values[x] = ds.Find(ds.values[x])
	return ds.values[x]
}

func (ds *DisjointSet) Union(x, y int) {
	a, b := ds.Find(x), ds.Find(y)

	if a != b {
		ds.values[b] = a
	}
}

func (ds *DisjointSet) GetValues() []int {
	return ds.values
}

func findCircleNum(isConnected [][]int) int {
	ds := Constructor(len(isConnected))

	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[0]); j++ {
			if i != j && isConnected[i][j] == 1 {
				ds.Union(i, j)
			}
		}
	}

	hm := make(map[int]int)
	values := ds.GetValues()
	for i := 0; i < len(values); i++ {
		ds.Find(i) // to compress path
		hm[values[i]]++
	}

	return len(hm)
}
