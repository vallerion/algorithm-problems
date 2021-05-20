package dsu

import "fmt"

// indexes: 	[1,2,3,5,8,9,10]
// values:		[1,1,1,5,5,5,10]
var valuesPC map[int]int

func findPC(i int) int {
	if valuesPC[i] == i {
		return i
	}

	valuesPC[i] = findPC(valuesPC[i])
	
	return valuesPC[i]
}

func makeSetPC(i int) {
	valuesPC[i] = i
}

func unionPC(a, b int) {
	as, bs := findPC(a), findPC(b)

	if as != bs {
		valuesPC[bs] = as
	}
}

func InitPC() {
	valuesPC = make(map[int]int, 0)
}

func MainDSU() {
	InitPC()

	for i := 1; i <= 8; i++ {
		makeSetPC(i)
	}

	fmt.Println(valuesPC)

	unionPC(1,2)
	unionPC(3,4)
	unionPC(5,6)
	unionPC(7,8)
	fmt.Println(valuesPC)

	unionPC(3,4) // already unioned
	unionPC(3,8)
	fmt.Println(valuesPC)
}