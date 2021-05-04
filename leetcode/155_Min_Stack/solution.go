package _155_Min_Stack

type defNode struct {
	value *minNode
	prev  *defNode
\

type stack struct {
	head *defNode
}

func (st *stack) push(mn *minNode) {
	dn := &defNode{mn, st.head}
	st.head = dn
}

func (st *stack) pop() {
	st.head = st.head.prev
}

func (st *stack) top() *defNode {
	return st.head
}

func (st *stack) isEmpty() bool {
	return st.head == nil
}


type minNode struct {
	value int
	prev  *minNode
}

type MinStack struct {
	top  *minNode
	mins stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return *new(MinStack)
}

func (this *MinStack) Push(val int) {
	nd := &minNode{val, this.top}
	this.top = nd

	if this.mins.isEmpty() || val < this.mins.top().value.value {
		this.mins.push(nd)
	}
}

func (this *MinStack) Pop() {
	if this.mins.top().value == this.top {
		this.mins.pop()
	}

	this.top = this.top.prev
}

func (this *MinStack) Top() int {
	return this.top.value
}

func (this *MinStack) GetMin() int {
	return this.mins.top().value.value
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
