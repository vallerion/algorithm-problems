package _399_Evaluate_Division

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

// Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	cnt := 0
	variables := make(map[string]int)

	for i := 0; i < len(equations); i++ {
		if _, ok := variables[equations[i][0]]; !ok {
			variables[equations[i][0]] = cnt
			cnt++
		}
		if _, ok := variables[equations[i][1]]; !ok {
			variables[equations[i][1]] = cnt
			cnt++
		}
	}

	ds := Constructor(cnt)
	edgeValues := make([][]float64, cnt)

	for i := 0; i < len(equations); i++ {
		ds.Union(variables[equations[i][0]], variables[equations[i][1]])

		if len(edgeValues[variables[equations[i][0]]]) == 0 {
			edgeValues[variables[equations[i][0]]] = make([]float64, cnt)
		}

		if len(edgeValues[variables[equations[i][1]]]) == 0 {
			edgeValues[variables[equations[i][1]]] = make([]float64, cnt)
		}

		edgeValues[variables[equations[i][0]]][variables[equations[i][1]]] = values[i]
		edgeValues[variables[equations[i][1]]][variables[equations[i][0]]] = 1 / values[i]
	}

	result := make([]float64, 0)
	for i := 0; i < len(queries); i++ {
		_, cExists := variables[queries[i][0]]
		_, dExists := variables[queries[i][1]]
		if !cExists || !dExists || !ds.IsConnected(variables[queries[i][0]], variables[queries[i][1]]) {
			result = append(result, -1.)
			continue
		}

		sum := dfs(variables[queries[i][0]], variables[queries[i][0]], variables[queries[i][1]], &edgeValues)
		result = append(result, sum)
	}

	return result
}

func dfs(parent, curr, target int, edgeValues *[][]float64) float64 {
	if curr == target {
		return 1.
	}
	if (*edgeValues)[curr][target] > 0 {
		return (*edgeValues)[curr][target]
	}

	for i := 0; i < len((*edgeValues)[curr]); i++ {
		if (*edgeValues)[curr][i] > 0 && i != parent {
			d := dfs(curr, i, target, edgeValues)
			if d > 0 {
				(*edgeValues)[curr][target] = (*edgeValues)[curr][i] * d
			}

		}
	}

	return (*edgeValues)[curr][target]
}
