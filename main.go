package main

import "fmt"
import . "./cormen/sorting"

func main() {
	fmt.Print(
		Merge([]int{1,30,50,2,25,90}, 0, 3, 6),
	)
}
