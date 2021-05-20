package dsu

// indexes: 	[1,2,3,5,8,9,10]
// values:		[1,1,1,5,5,5,10]
// ranks:		[0,0,0,0,0,0,0]
var valuesRank map[int]int
var ranks map[int]int

func findRank(i int) int {
	if valuesRank[i] == i {
		return i
	}
	
	return findRank(valuesRank[i])
}

func makeSetRank(i int) {
	valuesRank[i] = i
	ranks[i] = 0
}

func unionRank(a, b int) {
	as, bs := findRank(a), findRank(b)
	if as == bs {
		return
	}

	if ranks[as] > ranks[bs] {
		valuesRank[bs] = as

		if ranks[as] == ranks[bs] {
			ranks[as]++
		}
	} else {
		valuesRank[as] = bs

		if ranks[as] == ranks[bs] {
			ranks[bs]++
		}
	}
}

func InitRank() {
	valuesRank = make(map[int]int, 0)
	ranks = make(map[int]int, 0)
}