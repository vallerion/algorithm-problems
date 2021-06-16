package basic

import "fmt"
import "strconv"
import "strings"
import "bufio"
import "os"


func main() {
	in := bufio.NewReader(os.Stdin)

	in.ReadString('\n')
	input, _ := in.ReadString('\n')
	input = strings.TrimRight(input, "\n")

	var root int
	tb := make(map[int][]int, 0)

	for i, v := range strings.Split(input, " ") {
		node, _ := strconv.Atoi(v)
		if node == -1 {
			root = i
			continue
		}
		tb[node] = append(tb[node], i)
	}

	queue := tb[root]
	height := 0

	// fmt.Println(tb)
	// fmt.Println(queue)

	for len(queue) > 0 {
		curr := queue
		queue = []int{}

		for _, v := range curr {
			if _, ok := tb[v]; ok {
				queue = append(queue, tb[v]...)
			}
		}

		height++
	}

	fmt.Println(height+1)
}