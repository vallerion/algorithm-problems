package main

import "fmt"
import . "./sedvik/1_1"

func main() {
	fmt.Print(
		QuickFind(GeneratePairs(5)),
	)
}
