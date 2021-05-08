package main

import "fmt"
//import . "./cormen/exercises"
import . "./lessons"

func main() {
	heap := BuildMinHeap([]int{1,2,3,4,5})
	fmt.Println(heap)
}
