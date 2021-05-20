package dsu

// array with valuesNaive items from every set
// example:
// there are 3 sets: {1,2,3} {10} {5,9,8}
// we know values member in every set,
// so values will be like:
// indexes: 	[1,2,3,5,8,9,10]
// values:		[1,1,1,5,5,5,10]
var valuesNaive map[int]int

func findNaive(i int) int {
	if valuesNaive[i] == i {
		return i
	}
	
	return findNaive(valuesNaive[i])
}

func makeSetNaive(i int) {
	valuesNaive[i] = i
}

func unionNaive(a, b int) {
	as, bs := findNaive(a), findNaive(b)

	if as != bs {
		valuesNaive[bs] = as
	}
}

func InitNaive() {
	valuesNaive = make(map[int]int, 0)
}
