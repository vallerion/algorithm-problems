package roads_not_only_in_berland

import "fmt"

// DSU task
// https://codeforces.com/contest/25/problem/D

var roads map[int]int

func makeSet(v int) {
	//if roads[v] != 0 {
	//	return
	//}

	roads[v] = v
}

func find(v int) int {
	if roads[v] == v {
		return v
	}

	roads[v] = find(roads[v])

	return roads[v]
}

func union(a, b int) bool {
	av, bv := find(a), find(b)

	if av != bv {
		roads[bv] = av
		return true
	}

	return false
}

func main() {
	var cityCount int
	var merge, remove, input []int

	fmt.Scanf("%d\n", &cityCount)
	roads = make(map[int]int, cityCount)

	for i := 0; i < cityCount; i++ {
		makeSet(i)
	}

	for i := 0; i < cityCount-1; i++ {
		var a, b int
		fmt.Scanf("%d %d\n", &a, &b)
		input = append(input, a, b)

		u := union(a, b)
		if u == false {
			remove = append(remove, a, b)
		}
	}

	for i := 2; i <= cityCount; i++ {
		if find(1) != find(i) {
			merge = append(merge, 1, i)
			union(1, i)
		}
	}

	fmt.Println(len(merge) / 2)
	for i, j := 0, 0; i < len(merge) && j < len(remove); {
		fmt.Printf("%d %d %d %d\n", remove[j], remove[j+1], merge[i], merge[i+1])

		i += 2
		j += 2
	}
}
