//package hash_tables
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
 * Task #5.2
 * @see ./statements.pdf
 */

type node struct {
	val  string
	next *node
}

var hashmap map[int]*node
var m, n int

const p = 1000000007 // 1 000 000 007
const x = 263

func hash(input string) int {
	h := 0

	for i, s := range input {
		// https://brestprog.by/topics/modulo/
		h += int(s) * pow(x, i)
		//h += int(s) * int(math.Pow(x, float64(i)))
		//h += (int(s) % p) * (int(math.Pow(x, float64(i))) % p)
	}

	return (h % p) % m
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}

	res := a

	for b > 1 {
		res = (res * a) % p
		//res = ((res % p) * (a % p)) % p
		b--
	}

	return res
}

// возвращаем новый head
func delFromList(head *node, query string) *node {
	if head == nil {
		return nil
	}

	var prev *node
	curr := head

	for curr != nil {
		if curr.val == query {
			if prev != nil {
				// если есть предыдущий то сцепляем его и следующий теряя текущий
				prev.next = curr.next
			} else {
				// если предыдущего нет, значит head надо заменить на на следующий
				head = curr.next
			}
			break
		}

		prev = curr
		curr = curr.next
	}

	return head
}

func main() {
	fmt.Scanf("%d\n", &m)
	fmt.Scanf("%d\n", &n)

	hashmap = make(map[int]*node, m)

	if n == 0 {
		fmt.Println("")
		return
	}

	var action, value string

	scanner := bufio.NewScanner(os.Stdin)

	for n > 0 {
		n--
		scanner.Scan()

		fmt.Sscanf(scanner.Text(), "%s %s\n", &action, &value)
		handleAction(action, value)
	}
}

func handleAction(action, value string) {
	switch action {
	case "add":
		h := hash(value)
		if hashmap[h] == nil {
			hashmap[h] = &node{value, nil}
		} else {
			curr := hashmap[h]
			found := false
			for curr != nil {
				if curr.val == value {
					found = true
					break
				}

				curr = curr.next
			}

			if found == false {
				newHead := &node{value, hashmap[h]}
				hashmap[h] = newHead
			}
		}

		break
	case "del":
		h := hash(value)
		hashmap[h] = delFromList(hashmap[h], value)
		break
	case "find":
		h := hash(value)
		curr := hashmap[h]
		found := false
		for curr != nil {
			if curr.val == value {
				found = true
				break
			}

			curr = curr.next
		}

		if found {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}

		break
	case "check":
		v, _ := strconv.Atoi(value)
		curr, ok := hashmap[v]

		for curr != nil && ok {
			fmt.Print(curr.val)
			if curr.next != nil {
				fmt.Print(" ")
			}
			curr = curr.next
		}
		fmt.Print("\n")
		break
	}
}
