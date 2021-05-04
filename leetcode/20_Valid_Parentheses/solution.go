package _20_Valid_Parentheses

// '(', ')', '{', '}', '[' and ']'
func isValid(s string) bool {
	stk := new(stack)

	for _, sym := range s {
		if sym == '(' || sym == '{' || sym == '[' {
			stk.push(sym)
			continue
		}
		if stk.isEmpty() {
			return false
		}

		if sym == ')' && stk.pop().value != '(' {
			return false
		}
		if sym == '}' && stk.pop().value != '{' {
			return false
		}
		if sym == ']' && stk.pop().value != '[' {
			return false
		}
	}

	return stk.isEmpty()
}

type listNode struct {
	value rune
	prev  *listNode
}

type stack struct {
	top *listNode
}

func (st *stack) push(v rune) {
	ln := &listNode{v, st.top}
	st.top = ln
}

func (st *stack) pop() *listNode {
	temp := st.top
	st.top = st.top.prev
	return temp
}

func (st *stack) isEmpty() bool {
	return st.top == nil
}
