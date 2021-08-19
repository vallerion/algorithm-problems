package _1791_Find_Center_of_Star_Graph

func findCenter(edges [][]int) int {
	adjacencyList := make([][]int, len(edges)+2)

	for i := 0; i < len(edges); i++ {
		if adjacencyList[edges[i][0]] == nil {
			adjacencyList[edges[i][0]] = make([]int, 0)
		}
		if adjacencyList[edges[i][1]] == nil {
			adjacencyList[edges[i][1]] = make([]int, 0)
		}

		adjacencyList[edges[i][0]] = append(adjacencyList[edges[i][0]], edges[i][1])
		adjacencyList[edges[i][1]] = append(adjacencyList[edges[i][1]], edges[i][0])
	}

	maxVertex := 0
	maxLength := len(adjacencyList[maxVertex])
	for i := 1; i < len(adjacencyList); i++ {
		if len(adjacencyList[i]) > maxLength {
			maxLength = len(adjacencyList[i])
			maxVertex = i
		}
	}

	return maxVertex
}
