package may_contest

// [9,3,5] -> 9-5-3=1 -> [1,3,5]
// [1,3,5] -> 5-3-1=1 -> [1,3,1]
// [1,3,1] -> 3-1-1=1 -> [1,1,1] // sum == len -> true
//
// [1,1,1,2] -> 2-1-1-1=-1 < 0 -> false
//
// [8,5] -> 8-5=3 -> [3,5]
// [3,5] -> 5-3=2 -> [3,2]
// [3,2] -> 3-2=1 -> [1,2]
// [1,2] -> 2-1=1 -> [1,1] // sum == len -> true
//
// plan:
// 1. build heap to faster find max item, also calc sum on every action with heap
// 2. eject max
// 3. calc diff = max - (sum - max)
// 4. if diff <= 0 return false, if diff > 0 insert to heap
// 5. repeat it until sum != len
//
// difficult cases: [1,1000000000], [2,900000001], [2], [1]
// okay, lets change algo
//
// new plan:
// 1. build heap to faster find max item, also calc sum on every action with heap
// 2. eject max
// 3. calc diff = max%sum
// 4. if diff == 0 && sum == 1 -> diff = 1
// 5. if diff <= 0 || sum > max -> return false
//
// [2,900000001] -> 900000001%2=1 -> [2,1]
// [2,1] -> 2%1=0 -> diff=0&&sum==1 -> 1 -> [1,1]
//
// [1,1,2] -> 2%2=0 -> false
//
// [1,1,1,2] -> sum>max -> false
//
// [9,3,5] -> 9%8=1 -> [1,3,5]
// [1,3,5] -> 5%4=1 -> [1,3,1]
// [1,3,1] -> 3%2=1 -> [1,1,1]
//
// [1,1000000000] -> 1000000000%1=0 -> diff=0&&sum==1 -> 1 -> [1,1]

type maxHeap struct {
	values []int
	sum    int
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
	h.sum += value
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

	h.sum -= head
	return head
}

func (h *maxHeap) getMax() int {
	return h.values[0]
}

func (h *maxHeap) length() int {
	return len(h.values)
}

func buildHeap(nums []int) maxHeap {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	h := maxHeap{nums, sum}
	for i := len(nums) / 2; i >= 0; i-- {
		h.siftDown(i)
	}

	return h
}

func isPossible(target []int) bool {
	if len(target) == 1 {
		if target[0] == 1 {
			return true
		}
		return false
	}

	h := buildHeap(target)

	for h.sum != h.length() {
		first := h.ejectMax()
		diff := first % h.sum

		if diff == 0 && h.sum == 1 {
			diff = 1
		}

		if diff <= 0 || h.sum > first {
			return false
		}
		h.insert(diff)
	}

	return true
}
