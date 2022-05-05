package _225_Implement_Stack_using_Queues

// You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
// q = append(q, x) - push
// q[0] - peak
// q[0], q=q[1:] = pop

type MyStack struct {
	q1, q2 []int
}

func Constructor() MyStack {
	return MyStack{make([]int, 0), make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)
}

func (this *MyStack) Pop() int {
	for i := len(this.q1) - 1; i >= 0; i-- {
		this.q2 = append(this.q2, this.q1[i])
	}
	this.q1 = []int{}

	top := this.Top()
	this.q2 = this.q2[1:]

	return top
}

func (this *MyStack) Top() int {
	if len(this.q2) > 0 {
		return this.q2[0]
	}
	if len(this.q1) > 0 {
		return this.q1[len(this.q1)-1]
	}
	return -1
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0 && len(this.q2) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
