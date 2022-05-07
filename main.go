package main

import "fmt"

func main() {
	st := make([]int, 0)

	values := [...]int{4, 5, 6, 1, 2, 3, 94, 6, 2}

	for i := 0; i < len(values); i++ {

		for len(st) > 0 && st[len(st)-1] > values[i] {
			st = st[:len(st)-1]
		}

		st = append(st, values[i])
	}

	fmt.Println(st)
}
