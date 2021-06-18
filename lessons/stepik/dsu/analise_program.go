//package dsu
package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 * Task #4.2
 * @see ./statements.pdf
 */

func find(set map[int]int, x int) int {
	if set[x] == x {
		return x
	}

	set[x] = find(set, set[x])
	return set[x]
}

func union(set map[int]int, a, b int) bool {
	ax, bx := find(set, a), find(set, b)

	if ax == bx {
		return false
	}

	set[bx] = ax
	return true
}

func main() {
	var n, e, d int
	fmt.Scanf("%d %d %d\n", &n, &e, &d)

	set := make(map[int]int)
	var x, y int

	scanner := bufio.NewScanner(os.Stdin)

	for n > 0 {
		set[n] = n
		n--
	}

	for e > 0 {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)

		union(set, x, y)

		e--
	}

	correct := true

	for d > 0 {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)

		xx, yy := find(set, x), find(set, y)

		if xx == yy {
			correct = false
			break
		}

		d--
	}

	if correct {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
