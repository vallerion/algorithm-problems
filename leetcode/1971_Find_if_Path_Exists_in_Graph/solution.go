package _1971_Find_if_Path_Exists_in_Graph

func validPath(n int, edges [][]int, start int, end int) bool {
	if n == 1 {
		return true
	}

	adjacencyList := make([][]int, n)

	for i := 0; i < len(edges); i++ {
		adjacencyList[edges[i][0]] = append(adjacencyList[edges[i][0]], edges[i][1])
		adjacencyList[edges[i][1]] = append(adjacencyList[edges[i][1]], edges[i][0])
	}

	visited := make([]bool, n)
	q := make([]int, 0)
	q = append(q, adjacencyList[start]...)
	visited[start] = true
	isValid := false

	for len(q) > 0 && isValid == false {
		ln := len(q)
		for i := 0; i < ln; i++ {
			if q[i] == end {
				isValid = true
				break
			}

			visited[q[i]] = true

			for j := 0; j < len(adjacencyList[q[i]]); j++ {
				if visited[adjacencyList[q[i]][j]] == false {
					q = append(q, adjacencyList[q[i]][j])
				}
			}
		}

		q = q[ln:]
	}

	return isValid
}
