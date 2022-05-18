package _1192_Critical_Connections_in_a_Network

func criticalConnections(n int, connections [][]int) [][]int {
	adj := make([][]int, n)
	visited := make([]bool, n)
	result := make([][]int, 0)
	depths, low := make([]int, n), make([]int, n)

	for i := 0; i < len(connections); i++ {
		adj[connections[i][0]] = append(adj[connections[i][0]], connections[i][1])
		adj[connections[i][1]] = append(adj[connections[i][1]], connections[i][0])
	}

	dfs(0, 0, 0, &depths, &low, &visited, adj, &result)

	return result
}

func dfs(parent, curr, depth int, depths, low *[]int, visited *[]bool, adj [][]int, res *[][]int) {
	(*visited)[curr] = true
	(*depths)[curr] = depth
	(*low)[curr] = depth

	for i := 0; i < len(adj[curr]); i++ {
		if adj[curr][i] == parent {
			continue
		}

		if (*visited)[adj[curr][i]] == false {
			dfs(curr, adj[curr][i], depth+1, depths, low, visited, adj, res)
		}

		(*low)[curr] = min((*low)[curr], (*low)[adj[curr][i]])

		if (*low)[adj[curr][i]] > (*depths)[curr] {
			*res = append(*res, []int{adj[curr][i], curr})
		}
	}
}

// d=[0,0,0,0]
// l=[0,0,0,0]

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
