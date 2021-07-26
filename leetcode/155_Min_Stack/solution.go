package _155_Min_Stack

type MinStack struct {
	values, mins []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)

	if len(this.mins) == 0 || val <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	top := this.values[len(this.values)-1]
	this.values = this.values[:len(this.values)-1]

	if len(this.mins) > 0 && top == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}