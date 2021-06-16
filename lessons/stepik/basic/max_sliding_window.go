package basic
//package main

/*
 * Task #2.5
 * @see ./statements.pdf
 */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 2 7 3 1 5 2 6 2
// 4
//
// 1. наполняю 1й стек до 4х элементов
// 2. на 5м:
// 2.1 беру максимум на данный момент из двух стеков -> max(st1.max(), st2.max())
// 2.2 перекидываю все элементы во второй стек
// 2.3 делаю pop из второго стека
// 2.4 делаю push в первый стек
//
// st1 = [2 7 3 1] [2 7]
// st2 = [] []
// max = 7
// ...
// st1 = [5] [5]
// st2 = [1 3 7] [1 3 7]
// max = 7
// ...
// st1 = [5 2] [5]
// st2 = [1 3] [1 3]
// max = 5
// ...
// st1 = [5 2 6] [5 6]
// st2 = [1] [1]
// max = 6
// ...
// st1 = [5 2 6 2] [5 6]
// st2 = [1] [1]
// max = 6

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	input := scanner.Text()
	scanner.Scan()
	window, _ := strconv.Atoi(scanner.Text())

	if n == 0 {
		return
	}

	st1, st2, maxSt1, maxSt2 := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)

	splitedInput := strings.Split(input, " ")

	for j, v := range splitedInput {
		num, _ := strconv.Atoi(v)

		st1 = append(st1, num)
		if len(maxSt1) == 0 || st1[maxSt1[len(maxSt1)-1]] < num {
			maxSt1 = append(maxSt1, len(st1)-1)
		}

		if len(st2) > 0 {
			if maxSt2[len(maxSt2)-1] == len(st2)-1 {
				maxSt2 = maxSt2[:len(maxSt2)-1]
			}
			st2 = st2[:len(st2)-1]
		}

		if len(st1) == window {
			for i := len(st1) - 1; i >= 0; i-- {
				st2 = append(st2, st1[i])

				if len(maxSt2) == 0 || st2[maxSt2[len(maxSt2)-1]] < st1[i] {
					maxSt2 = append(maxSt2, len(st2)-1)
				}
			}
			st1, maxSt1 = make([]int, 0), make([]int, 0)
		}

		if j < window-1 {
			continue
		}

		if len(st1) == 0 {
			fmt.Print(st2[maxSt2[len(maxSt2)-1]])
		} else if len(st2) == 0 {
			fmt.Print(st1[maxSt1[len(maxSt1)-1]])
		} else {
			fmt.Print(max(st1[maxSt1[len(maxSt1)-1]], st2[maxSt2[len(maxSt2)-1]]))
		}
		if j < len(splitedInput)-1 {
			fmt.Print(" ")
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
