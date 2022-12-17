package _150_Evaluate_Reverse_Polish_Notation

import "strconv"

func evalRPN(tokens []string) int {
	st := make([]int, 0)

	for i := 0; i < len(tokens); i++ {
		if isSign(tokens[i]) {
			calc := calculate(st[len(st)-2], st[len(st)-1], tokens[i])
			st = st[:len(st)-2]
			st = append(st, calc)
		} else {
			digit, _ := strconv.Atoi(tokens[i])
			st = append(st, digit)
		}
	}

	return st[0]
}

func calculate(first, second int, sign string) int {
	switch sign {
	case "+":
		return first + second
	case "-":
		return first - second
	case "/":
		return first / second
	case "*":
		return first * second
	}

	return 0
}

func isSign(symbol string) bool {
	switch symbol {
	case "+", "-", "/", "*":
		return true
	}
	return false
}
