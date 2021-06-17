//package hash_tables
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Task #4.1
 * @see ./statements.pdf
 */

func find(set map[int]int, x int) int {
	if set[x] <= 0 {
		return x
	}

	set[x] = find(set, set[x])
	return set[x]
}

func union(set map[int]int, a, b int) int {
	ax, bx := find(set, a), find(set, b)

	if ax == bx {
		return 0
	}

	set[ax] = -(absInt(set[ax]) + absInt(set[bx]))
	set[bx] = ax

	return absInt(set[ax])
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n, m, highest int
	fmt.Scanf("%d %d\n", &n, &m)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")

	tables := make(map[int]int)

	for i, s := range strings.Split(input, " ") {
		si, _ := strconv.Atoi(s)
		tables[i+1] = -si

		highest = max(highest, si)
	}

	//fmt.Println(tables)

	var destination, source, height int

	//m = 1
	for m > 0 {
		input, _ = reader.ReadString('\n')
		fmt.Sscanf(input, "%d %d", &destination, &source)
		//fmt.Println(destination, source)

		height = union(tables, destination, source)
		highest = max(height, highest)
		fmt.Println(highest)

		m--
	}
}
