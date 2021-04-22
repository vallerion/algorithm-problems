package __1

import (
	"math/rand"
	"time"
)

const N = 10

var PairExample1 = [][]int{{0, 2}, {1, 4}, {2, 5}, {3, 6}, {0, 4}, {6, 0}, {1, 3}}
var PairExample2 = [][]int{
	{3, 4}, {4, 9}, {8, 0}, {2, 3}, {5, 6}, {2, 9},
	{5, 9}, {7, 3}, {4, 8}, {5, 6}, {0, 2}, {6, 1},
}

func GeneratePairs(length int) [][]int {
	result := make([][]int, 0)

	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		result = append(result, []int{rand.Intn(10), rand.Intn(10)})
	}

	return result
}

// QuickFind
// ex. input: [[0,2],[1,4]...]
func QuickFind(pairs [][]int) [][]int {
	result := make([][]int, 0)
	id := make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = i
	}

	for pi := 0; pi < len(pairs); pi++ {
		p := pairs[pi][0]
		q := pairs[pi][1]
		t := id[p]

		if t == id[q] {
			continue
		}

		for i := 0; i < N; i++ {
			if id[i] == t {
				id[i] = id[q]
			}
		}

		result = append(result, []int{p, q})
	}

	return result
}
