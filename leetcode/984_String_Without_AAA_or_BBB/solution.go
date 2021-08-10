package _984_String_Without_AAA_or_BBB

// 1 4

// biggest / (smallest+1) -> 4 / (1+1) -> 2 (every 2 place smallest)
// b b a b b

// 2 5
// 5 / 3 + 1 -> 1 + 1 -> 2
// 0 1 2 3 4 5 6 7 8

func strWithout3a3b(a int, b int) string {
	result := ""

	for a > 0 || b > 0 {
		writeA := false

		if len(result) > 1 && result[len(result)-1] == result[len(result)-2] {
			if result[len(result)-1] == 'b' {
				writeA = true
			}
		} else {
			if a > b {
				writeA = true
			}
		}

		if writeA == true {
			result += "a"
			a--
		} else {
			result += "b"
			b--
		}

	}

	return result
}
