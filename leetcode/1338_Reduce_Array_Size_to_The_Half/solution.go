package _1338_Reduce_Array_Size_to_The_Half

var hm map[int]int

func add(heap *[]int, value int) {
	*heap = append(*heap, value)
	siftUp(heap, len(*heap)-1)
}

func siftUp(heap *[]int, index int) {
	if index == 0 {
		return
	}

	curr, parent := (*heap)[index], (*heap)[(index-1)/2]

	for hm[curr] > hm[parent] {
		swap(heap, index, (index-1)/2)
		index = (index - 1) / 2

		curr, parent = (*heap)[index], (*heap)[(index-1)/2]
	}
}

func ejectMax(heap *[]int) int {
	max := (*heap)[0]
	(*heap)[0] = (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]

	siftDown(heap, 0)

	return max
}

func siftDown(heap *[]int, index int) {
	for index*2+1 < len(*heap) {
		leftIndex, rightIndex := index*2+1, index*2+2

		childToSwap := leftIndex

		if rightIndex < len(*heap) && hm[(*heap)[rightIndex]] > hm[(*heap)[leftIndex]] {
			childToSwap = rightIndex
		}

		if hm[(*heap)[index]] >= hm[(*heap)[childToSwap]] {
			break
		}

		swap(heap, index, childToSwap)
		index = childToSwap
	}
}

func swap(heap *[]int, x, y int) {
	temp := (*heap)[x]
	(*heap)[x] = (*heap)[y]
	(*heap)[y] = temp
}

func minSetSize(arr []int) int {
	hm = make(map[int]int)
	heap := make([]int, 0)
	nums := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		hm[arr[i]]++

		if hm[arr[i]] == 1 {
			nums = append(nums, arr[i])
		}
	}

	for i := 0; i < len(nums); i++ {
		add(&heap, nums[i])
	}

	ejectedMax := ejectMax(&heap)
	total := hm[ejectedMax]
	count := 1

	for len(heap) > 0 && total < len(arr)/2 {
		ejectedMax = ejectMax(&heap)
		total += hm[ejectedMax]
		count++
	}

	return count
}
