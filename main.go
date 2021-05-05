package main

import "fmt"
import . "./cormen/exercises"

func main() {
	//fmt.Print(
	//	MergeSort([]int{1,4,4,4,7,7,30,50,56,56,56,57,64,67,356,456,457,567,856}),
	//)
	fmt.Print(
		Bsi([]int{1,4,4,4,7,7,30,50,56,56,56,57,64,67,356,456,457,567,856}, 0, 18, 457),
	)
}
