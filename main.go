package main

import "fmt"
import . "./cormen/sorting"

func main() {
	fmt.Print(
		InsertionSort([]int{1,43,4,45,6,7,4,3,2}),
	)
}
