package main

import "fmt"
import . "./cormen/exercises"

func main() {
	fmt.Println(
		InsertionSort([]int{2,4,4,1,7,7,30,50,56,56,56,57,64,67,356,456,457,567,856}),
	)
	fmt.Println(
		InsertionSortBS([]int{2,4,4,1,7,7,30,50,56,56,56,57,64,67,356,456,457,567,856}),
	)
}
