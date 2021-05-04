package _239_Sliding_Window_Maximum

type dqueue struct {
	values []int
}

func (q *dqueue) pushFront(val int) {
	q.values = append([]int{val}, q.values...)
}

func (q *dqueue) pushBack(val int) {
	q.values = append(q.values, val)
}

func (q *dqueue) front() int {
	return q.values[0]
}

func (q *dqueue) back() int {
	return q.values[len(q.values)-1]
}

func (q *dqueue) popFront() {
	q.values = q.values[1:]
}

func (q *dqueue) popBack() {
	q.values = q.values[:len(q.values)-1]
}

func (q *dqueue) length() int {
	return len(q.values)
}

// 1,3,-1,-3,5,3,1,1
//
// i=0, j=0 [1]
// i=0, j=1 nums[j] > back -> popBack -> [] -> [3]
// i=0, j=2 nums[j] > back -> [3, -1] // 3
// i=1, j=3 nums[j] > back -> [3,-3]
//
// if j-i+1 < k -> j++
// else {
// 	result.push(3)
// 	if front == nums[i]
// 		popFront
// 	i++
// 	j++
// }
//
// [1,3,1,2,0,5]
// 3
// [1] -> [3] -> [3,1] // 3
// [3,2] // 3
// [2] -> [2,0] // 2
// [5] // 5

func maxSlidingWindow(nums []int, k int) []int {
	q := dqueue{[]int{}}
	maxs := make([]int, len(nums)-(k-1))

	i := 0
	j := 0

	for j < len(nums) {
		// for q.length() > 0 && nums[j] > q.front() {
		// 	q.popFront()
		// }
		for q.length() > 0 && nums[j] > q.back() {
			q.popBack()
		}
		q.pushBack(nums[j])

		if j-i+1<k {
			j++
		} else {
			maxs[i] = q.front()
			if nums[i] == q.front() {
				q.popFront()
			}

			i++
			j++
		}
	}

	return maxs
}