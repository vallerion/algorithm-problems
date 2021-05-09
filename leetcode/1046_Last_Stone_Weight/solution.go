package _1046_Last_Stone_Weight

// Plan:
// 1 build max-heap from array
// 2 eject max
// 3 eject max again
// 4 calc difference between two max
// 5 if difference more than 0 -> insert to heap
// 6 if heap.size > 1 go to step 2
// So i need methods:
// - build
// - insert
// - ejectMax
// Auxiliary methods:
// - parent
// - leftChild
// - rightChild
// - siftUp
// - siftDown

type maxHeap struct {
	values []int
}

func (h *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *maxHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *maxHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *maxHeap) insert(value int) {
	h.values = append(h.values, value)
	h.siftUp(len(h.values) - 1)
}

func (h *maxHeap) siftDown(i int) {
	for {
		left := h.leftChild(i)
		right := h.rightChild(i)
		childToSwap := i

		if left >= len(h.values) && right >= len(h.values) {
			break
		}

		if left >= len(h.values) && right < len(h.values) {
			childToSwap = right
		} else if right >= len(h.values) && left < len(h.values) {
			childToSwap = left
		} else if h.values[left] >= h.values[right] {
			childToSwap = left
		} else if h.values[right] > h.values[left] {
			childToSwap = right
		}

		if h.values[childToSwap] <= h.values[i] {
			break
		}

		h.values[childToSwap], h.values[i] = h.values[i], h.values[childToSwap]
		i = childToSwap
	}
}

func (h *maxHeap) siftUp(i int) {
	for {
		parent := h.parent(i)

		if h.values[parent] < h.values[i] {
			h.values[parent], h.values[i] = h.values[i], h.values[parent]
			i = h.parent(i)
		} else {
			break
		}
	}
}

func (h *maxHeap) ejectMax() int {
	head := h.values[0]
	h.values[0] = h.values[len(h.values)-1]
	h.values = h.values[:len(h.values)-1]
	h.siftDown(0)

	return head
}

func (h *maxHeap) getMax() int {
	return h.values[0]
}

func (h *maxHeap) length() int {
	return len(h.values)
}

func buildHeap(nums []int) maxHeap {
	h := maxHeap{nums}

	for i := len(nums) / 2; i >= 0; i-- {
		h.siftDown(i)
	}

	return h
}

func lastStoneWeight(stones []int) int {
	h := buildHeap(stones)

	for h.length() > 1 {
		first := h.ejectMax()
		second := h.ejectMax()
		diff := absInt(first - second)

		if diff > 0 {
			h.insert(diff)
		}
	}

	if h.length() == 0 {
		return 0
	}

	return h.getMax()
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
