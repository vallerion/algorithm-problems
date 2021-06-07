package stepik
//package main

/*
 * Task #2.4
 * @see ./statements.pdf
 */

import (
	"bufio"
	"fmt"
	"os"
)

var stack, maxIndexes []int

func main() {
	var n int
	stack, maxIndexes = make([]int, 0), make([]int, 0)

	fmt.Scanf("%d\n", &n)

	if n == 0 {
		fmt.Println("")
		return
	}

	var value int
	var action string

	scanner := bufio.NewScanner(os.Stdin)

	for n > 0 {
		n--
		scanner.Scan()

		fmt.Sscanf(scanner.Text(), "%s %d\n", &action, &value)
		handleAction(action, value)
	}
}

func handleAction(action string, value int) {
	switch action {
	case "push":
		stack = append(stack, value)
		if len(maxIndexes) == 0 {
			maxIndexes = append(maxIndexes, 0)
		} else if stack[maxIndexes[len(maxIndexes)-1]] < value {
			maxIndexes = append(maxIndexes, len(stack)-1)
		}
		break
	case "pop":
		if maxIndexes[len(maxIndexes)-1] == len(stack)-1 {
			maxIndexes = maxIndexes[:len(maxIndexes)-1]
		}
		stack = stack[:len(stack)-1]
		break
	case "max":
		fmt.Println(stack[maxIndexes[len(maxIndexes)-1]])
	}
}
