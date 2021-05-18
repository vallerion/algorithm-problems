package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// All tasks: https://stepik.org/media/attachments/lesson/41233/statements.pdf
// this 3.1

type MinHeap struct {
	values    []int
	swaps     int
	swapPairs []int
}

func BuildMinHeap(nums []int) MinHeap {
	mh := MinHeap{nums, 0, []int{}}
	for i := len(nums) / 2; i >= 0; i-- {
		mh.SiftDown(i)
	}
	return mh
}

func (h *MinHeap) Parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) LeftChild(i int) int {
	return i*2 + 1
}

func (h *MinHeap) RightChild(i int) int {
	return i*2 + 2
}

func (h *MinHeap) SiftUp(i int) {
	for parent := h.Parent(i); h.values[parent] > h.values[i]; parent = h.Parent(i) {
		h.values[parent], h.values[i] = h.values[i], h.values[parent]
		i = h.Parent(i)
	}
}

func (h *MinHeap) SiftDown(i int) {
	for {
		left := h.LeftChild(i)
		right := h.RightChild(i)
		childToSwap := i

		if left < len(h.values) {
			childToSwap = left
		}

		if right < len(h.values) {
			if left >= len(h.values) {
				childToSwap = right
			} else if h.values[right] < h.values[left] {
				childToSwap = right
			}
		}

		if h.values[i] <= h.values[childToSwap] {
			break
		}

		h.values[childToSwap], h.values[i] = h.values[i], h.values[childToSwap]
		h.swapPairs = append(h.swapPairs, i, childToSwap)
		h.swaps++

		i = childToSwap
	}
}

func (h *MinHeap) Insert(value int) {
	h.values = append(h.values, value)
	h.SiftUp(len(h.values) - 1)
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n') // first line
	secondLine, _ := reader.ReadString('\n')
	secondLine = strings.TrimRight(secondLine, "\n")

	strs := strings.Split(secondLine, " ")
	nums := strToInt(strs)

	heap := BuildMinHeap(nums)
	fmt.Println(heap.swaps)

	for i := 0; i < len(heap.swapPairs)-1; i+=2 {
		fmt.Printf("%v %v\n", heap.swapPairs[i], heap.swapPairs[i+1])
	}
}


func strToInt(input []string) []int {
	result := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		result[i], _ = strconv.Atoi(input[i])
	}

	return result
}