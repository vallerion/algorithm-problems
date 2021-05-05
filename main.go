package main

import "fmt"
import . "./cormen/sorting"

func main() {
	fmt.Print(
		Merge([]int{1,30,50}, []int{2,25,90}),
	)
	fmt.Print(
		MergeSort([]int{1,2,5,3,45,3,6,3,6,3,4,64,7,56,7,3}),
	)
}
