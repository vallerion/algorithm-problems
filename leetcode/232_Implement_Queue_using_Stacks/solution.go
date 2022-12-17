package _232_Implement_Queue_using_Stacks

type MyQueue struct {
	st1, st2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		st1: make([]int, 0),
		st2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.st1 = append(this.st1, x)
}

func (this *MyQueue) Pop() int {
	this.pour()
	top := this.st2[len(this.st2)-1]
	this.st2 = this.st2[:len(this.st2)-1]
	return top
}

func (this *MyQueue) Peek() int {
	this.pour()
	return this.st2[len(this.st2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.st1) == 0 && len(this.st2) == 0
}

func (this *MyQueue) pour() {
	if len(this.st2) > 0 {
		return
	}

	for len(this.st1) > 0 {
		top := this.st1[len(this.st1)-1]
		this.st1 = this.st1[:len(this.st1)-1]
		this.st2 = append(this.st2, top)
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
