package dsu

// indexes: 	[1,2,3,5,8,9,10]
// values:		[1,1,1,5,5,5,10]
// ranks:		[0,0,0,0,0,0,0]
var valuesPCRank map[int]int
var pcRanks map[int]int

func findPCRank(i int) int {
	if valuesPCRank[i] == i {
		return i
	}

	valuesPCRank[i] = findPCRank(valuesPCRank[i])

	return valuesPCRank[i]
}

func makeSetPCRank(i int) {
	valuesPCRank[i] = i
	pcRanks[i] = 0
}

func unionPCRank(a, b int) {
	as, bs := findPCRank(a), findPCRank(b)
	if as == bs {
		return
	}

	if pcRanks[as] > pcRanks[bs] {
		valuesPCRank[bs] = as

		if pcRanks[as] == pcRanks[bs] {
			pcRanks[as]++
		}
	} else {
		valuesPCRank[as] = bs

		if pcRanks[as] == pcRanks[bs] {
			pcRanks[bs]++
		}
	}
}

func InitPCRank() {
	valuesPCRank = make(map[int]int, 0)
	pcRanks = make(map[int]int, 0)
}
