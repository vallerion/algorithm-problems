package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 2
// 5 2
// LIZZIE KEVIN BOHDAN LALIT RUOYU
// 1 2
// 1 3
// 2 2
// KICK START
// 1 2
// 1 2

func main() {
	consoleScanner := bufio.NewScanner(os.Stdin)

	var caseCount int
	consoleScanner.Scan()
	fmt.Sscanf(consoleScanner.Text(), "%d\n", &caseCount)

	output := ""

	for i := 0; i < caseCount; i++ {
		output += fmt.Sprintf("Case #%d: ", i+1)

		var nameCount, queryCount int
		consoleScanner.Scan()
		fmt.Sscanf(consoleScanner.Text(), "%d %d\n", &nameCount, &queryCount)

		consoleScanner.Scan()
		namesStr := consoleScanner.Text()

		names := strings.Split(namesStr, " ")
		for j := 0; j < queryCount; j++ {
			var personA, personB int
			consoleScanner.Scan()
			fmt.Sscanf(consoleScanner.Text(), "%d %d\n", &personA, &personB)

			if canBeFriend(names[personA-1], names[personB-1]) {
				output += strconv.Itoa(2) + " "
			} else {
				steps := friendRange(names, personA-1, personB-1)
				output += strconv.Itoa(steps) + " "
			}
		}

		output += "\n"
	}

	fmt.Print(output)
}

func friendRange(names []string, personA, personB int) int {
	queueA, queueB := make([]int, 0), make([]int, 0)
	queueA, queueB = append(queueA, personA), append(queueB, personB)

	visitedA, visitedB := make([]bool, len(names)), make([]bool, len(names))
	visitedA[personA] = true
	visitedB[personB] = true

	steps := 0

	for len(queueA) > 0 || len(queueB) > 0 {
		nA, nB := len(queueA), len(queueB)

		if visitsIntersection(visitedA, visitedB) {
			return steps + 2
		}

		for i := 0; i < nA; i++ {
			for j := 0; j < len(names); j++ {
				if visitedA[j] == false && canBeFriend(names[queueA[i]], names[j]) {
					if j == personB {
						return steps + 2
					}

					visitedA[j] = true
					queueA = append(queueA, j)
				}
			}
		}

		for i := 0; i < nB; i++ {
			for j := 0; j < len(names); j++ {
				if visitedB[j] == false && canBeFriend(names[queueB[i]], names[j]) {
					if j == personA {
						return steps + 2
					}

					visitedB[j] = true
					queueB = append(queueB, j)
				}
			}
		}

		steps++
		queueA, queueB = queueA[nA:], queueB[nB:]
	}

	return -1
}

func visitsIntersection(visitsA, visitsB []bool) bool {
	for i := 0; i < len(visitsA); i++ {
		if visitsA[i] == true && visitsA[i] == visitsB[i] {
			return true
		}
	}

	return false
}

func canBeFriend(a, b string) bool {
	hp := make(map[rune]bool)

	for _, sym := range a {
		hp[sym] = true
	}

	for _, sym := range b {
		if _, exists := hp[sym]; exists {
			return true
		}
	}

	return false
}
