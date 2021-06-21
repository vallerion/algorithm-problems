package _20_Valid_Parentheses

func isValid(s string) bool {
	stack := make([]rune, 0)
	//symbols := map[string]string

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] == '(' && v == ')' {
			stack = stack[:len(stack)-1]
		} else if stack[len(stack)-1] == '{' && v == '}' {
			stack = stack[:len(stack)-1]
		} else if stack[len(stack)-1] == '[' && v == ']' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
