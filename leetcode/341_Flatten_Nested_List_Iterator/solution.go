package _341_Flatten_Nested_List_Iterator

type NestedInteger struct {
	list      []*NestedInteger
	value     int
	isInteger bool
}

// IsInteger Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this *NestedInteger) IsInteger() bool {
	return this.isInteger
}

// GetInteger Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this *NestedInteger) GetInteger() int {
	return this.value
}

// SetInteger Set this NestedInteger to hold a single integer.
func (this *NestedInteger) SetInteger(value int) {
	this.value = value
}

// Add Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {
	if this.list == nil {
		this.list = make([]*NestedInteger, 0)
	}

	this.list = append(this.list, &elem)
}

// GetList Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return this.list
}

// [[1,1],2,[1,1]]
// [[1,2],3] -> [1,2]
// 1 2 3

type Cursor struct {
	list  []*NestedInteger
	index int
}

type NestedIterator struct {
	cursorStack []*Cursor
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[]*Cursor{{nestedList, 0}}}
}

func (this *NestedIterator) Next() int {
	if this.HasNext() == false {
		return -1
	}

	for len(this.cursorStack) > 0 {
		curr := this.cursorStack[len(this.cursorStack)-1]

		if curr.index < len(curr.list) {
			if curr.list[curr.index].IsInteger() {
				result := curr.list[curr.index].GetInteger()
				curr.index++
				return result
			} else {
				this.cursorStack = append(this.cursorStack, &Cursor{curr.list[curr.index].GetList(), 0})
				curr.index++
			}
		} else {
			this.cursorStack = this.cursorStack[:len(this.cursorStack)-1]
		}
	}

	return -1
}

// [1,[1]] -> [1]

func (this *NestedIterator) HasNext() bool {
	if len(this.cursorStack) == 0 {
		return false
	}

	for len(this.cursorStack) > 0 {
		curr := this.cursorStack[len(this.cursorStack)-1]

		if curr.index < len(curr.list) {
			if curr.list[curr.index].IsInteger() {
				return true
			} else {
				this.cursorStack = append(this.cursorStack, &Cursor{curr.list[curr.index].GetList(), 0})
				curr.index++
			}
		} else {
			this.cursorStack = this.cursorStack[:len(this.cursorStack)-1]
		}
	}

	return false
}
