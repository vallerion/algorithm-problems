package main

import "fmt"
import . "./cormen/sorting"

func main() {
	fmt.Print(
		Merge([]int{3,5,1,2,25,90}, 0, 3, 6),
	)
}
