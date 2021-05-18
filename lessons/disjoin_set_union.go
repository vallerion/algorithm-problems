package lessons

import "fmt"

func find(smallest map[int]int, i int) int {
	return smallest[i]
}

func makeSet(smallest map[int]int, i int) {
	smallest[i] = i
}

func union(smallest map[int]int, a, b int) {
	as, bs := find(smallest, a), find(smallest, b)

	if as == bs {
		return
	}

	for i, v := range smallest {
		if v == bs {
			smallest[i] = as
		}
	}
}

// MainDSU
// array with smallest items from every set
// example:
// there are 3 sets: {1,2,3} {10} {5,9,8}
// we know smallest member in every set,
// so smallest will be like:
// indexes: 	[1,2,3,5,8,9,10]
// values:		[1,1,1,5,5,5,10]
func MainDSU() {
	smallest := make(map[int]int)

	makeSet(smallest, 1)
	makeSet(smallest, 2)
	makeSet(smallest, 3)
	makeSet(smallest, 5)
	makeSet(smallest, 8)
	makeSet(smallest, 9)
	makeSet(smallest, 10)
	fmt.Println(smallest)

	// {1,2,3}
	union(smallest, 1, 2)
	union(smallest, 2, 3)
	fmt.Println(smallest)
	// {5,9,8}
	union(smallest, 5, 9)
	union(smallest, 9, 8)
	fmt.Println(smallest)
}
