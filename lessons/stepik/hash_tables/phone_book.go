package hash_tables
//package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/*
 * Task #5.1
 * @see ./statements.pdf
 */

// 				 [number]name
var phoneBook map[int]string

func main() {
	var n int
	phoneBook = make(map[int]string, int(math.Pow(10, 5)))

	fmt.Scanf("%d\n", &n)

	if n == 0 {
		fmt.Println("")
		return
	}

	var number int
	var action, name string

	scanner := bufio.NewScanner(os.Stdin)

	for n > 0 {
		n--
		scanner.Scan()

		fmt.Sscanf(scanner.Text(), "%s %d %s\n", &action, &number, &name)
		handleAction(action, name, number)
	}
}

func handleAction(action, name string, number int) {
	switch action {
	case "add":
		phoneBook[number] = name
		break
	case "del":
		delete(phoneBook, number)
		break
	case "find":
		v, ok := phoneBook[number]
		if ok {
			fmt.Println(v)
		} else {
			fmt.Println("not found")
		}
	}
}
