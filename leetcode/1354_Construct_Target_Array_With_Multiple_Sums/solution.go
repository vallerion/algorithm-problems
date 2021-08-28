package _1354_Construct_Target_Array_With_Multiple_Sums

/**
[9,3,5]
[1,1,1] -> [1,3,1] -> [1,3,5] -> [9,3,5] -> true

[9,3,5] -> [3,5,1] -> [3,2,1] -> [1,1,2]
[9,3,5] -> 1,[3,5] -> [3,2] -> 1,[2] -> if len(A) == 1 && ( A[0] == len(source) || A[0] == len(source)-1)

[9,3,5] -> 9-(5+3)=1 -> 1,[3,5]
[3,5]   -> 5-3=2     -> [3,2]
[3,2]   -> 3-2=1     -> 1,[2]
[2]     -> if len(A) == 1 && ( A[0] == len(source) || A[0] == len(source)-1) -> true

---
[1,2,1,1]

[1,2,1,1] -> 2-3=-1 -> if < 0 -> false

---
[8,5]

[8,5] -> 8-5=3 -> [3,5]
[3,5] -> 5-3=2 -> [3,2]
[3,2] -> 3-2=1 -> 1,[2]
[2] -> if len(A) == 1 && ( A[0] == len(source) || A[0] == len(source)-1) -> true

[8,5] -> 8%5=3 -> [3,5]
[3,5] -> 5%3=2 -> [3,2]
[3,2] -> 3%2=1 -> [2]
[2] -> if len(A) == 1 && ( A[0] == len(source) || A[0] == len(source)-1) -> true

---
[9,3,4]

[9,3,4] -> 9-7=3 -> [3,3,4]
[3,3,4] -> 4-6=-2 -> if < 0 -> false

[9,3,4] -> 9%7=3 -> [3,3,4]
[3,3,4] -> 4<6 -> false

---
[1,1]

if max == 1 -> true

---
[1,2]

[1,2] -> 2-1=1 -> 1,[1]
[1] -> if len(A) == 1 && ( A[0] == len(source) || A[0] == len(source)-1) -> true

---
[1,1,2]

[1,1,2] -> 2%2=0 -> [1,1]
[1,1] -> 1%1=0 -> [1]
[1] -> if len(A) == 1 && ( A[0] == len(source) || A[0] == len(source)-1) -> false

---
[1,1,999999999]

[1,1,999999999] -> 999999999%2=1 -> [1,1]
[1,1] -> 1%1=0


*/

type MaxHeap struct {
	values []int
}

func Constructor() *MaxHeap {
	return &MaxHeap{values: make([]int, 0)}
}

// childLeft = i*2+1
// childRight = i*2+2
// parent = (i-1)/2

func (h *MaxHeap) build(source []int) {
	for i := 0; i < len(source); i++ {
		h.insert(source[i])
	}
}

func (h *MaxHeap) insert(value int) {
	h.values = append(h.values, value)
	h.siftUp(len(h.values) - 1)
}

func (h *MaxHeap) siftUp(index int) {
	if index <= 0 || index >= len(h.values) {
		return
	}

	parent := (index - 1) / 2

	for h.values[parent] < h.values[index] {
		h.swap(parent, index)

		index = parent
		parent = (index - 1) / 2
	}
}

func (h *MaxHeap) swap(one, two int) {
	h.values[one], h.values[two] = h.values[two], h.values[one]
}

func (h *MaxHeap) siftDown(index int) {
	if index < 0 || index >= len(h.values) {
		return
	}

	for {
		leftChild := index*2 + 1
		rightChild := index*2 + 2

		childToSwap := h.biggestIndex(leftChild, rightChild)

		if childToSwap == -1 || h.values[index] >= h.values[childToSwap] {
			break
		}

		h.swap(index, childToSwap)
		index = childToSwap
	}
}

func (h *MaxHeap) biggestIndex(one, two int) int {
	if one < len(h.values) && two < len(h.values) {
		if h.values[one] > h.values[two] {
			return one
		}
		return two
	}

	if one < len(h.values) {
		return one
	}

	if two < len(h.values) {
		return two
	}

	return -1
}

func (h *MaxHeap) extractMax() int {
	temp := h.values[0]
	h.values[0] = h.values[len(h.values)-1]
	h.values = h.values[:len(h.values)-1]

	h.siftDown(0)

	return temp
}

func (h *MaxHeap) getMax() int {
	return h.values[0]
}

func (h *MaxHeap) len() int {
	return len(h.values)
}

func isPossible(target []int) bool {
	if len(target) == 1 {
		if target[0] > 1 {
			return false
		} else {
			return true
		}
	}

	heap := Constructor()
	heap.build(target)

	sum := 0
	for i := 0; i < len(target); i++ {
		sum += target[i]
	}

	if heap.getMax() == 1 {
		return true
	}

	for heap.len() > 0 {
		max := heap.extractMax()

		sum -= max
		result := max % sum

		// cases: [1,1,999999999], [1,1,2]
		if result == 0 && sum == 1 {
			result = 1
		}

		if result <= 0 || max < sum {
			return false
		}

		sum += result
		heap.insert(result)
	}

	return true
}
