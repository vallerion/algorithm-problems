package _1338_Reduce_Array_Size_to_The_Half

var sizes map[int]int
var heap []int

func add(value int) {
	heap = append(heap, value)
	siftUp(len(heap) - 1)
}

func ejectTop() int {
	top := heap[0]
	heap[0] = heap[len(heap)-1]
	heap = heap[:len(heap)-1]

	siftDown(0)

	return top
}

func siftUp(index int) {
	if index < 1 {
		return
	}

	parent := (index - 1) / 2

	for sizes[heap[parent]] < sizes[heap[index]] {
		heap[parent], heap[index] = heap[index], heap[parent]

		index = parent
		parent = (index - 1) / 2
	}
}

func siftDown(index int) {
	for index*2+1 < len(heap) {
		left, right := index*2+1, index*2+2

		childToSwap := left

		if right < len(heap) && sizes[heap[right]] > sizes[heap[left]] {
			childToSwap = right
		}

		if sizes[heap[index]] >= sizes[heap[childToSwap]] {
			break
		}

		heap[childToSwap], heap[index] = heap[index], heap[childToSwap]
		index = childToSwap
	}
}

func minSetSize(arr []int) int {
	sizes = make(map[int]int)
	heap = make([]int, 0)

	for i := 0; i < len(arr); i++ {
		sizes[arr[i]]++
	}

	sum := 0
	for key, value := range sizes {
		add(key)
		sum += value
	}

	counter := 0
	for len(heap) > 0 && sum > len(arr)/2 {
		top := ejectTop()
		counter++
		sum -= sizes[top]
	}

	return counter
}
