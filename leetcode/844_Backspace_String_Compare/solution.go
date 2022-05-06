package _844_Backspace_String_Compare

// time - O(n)
func backspaceCompare(s string, t string) bool {
	sr, tr := []rune(s), []rune(t)

	backspaces := 0

	for i := len(sr) - 1; i >= 0; i-- {
		if sr[i] == '#' {
			backspaces++
			sr[i] = ' '
			continue
		}

		if backspaces > 0 {
			sr[i] = ' '
			backspaces--
		}
	}

	backspaces = 0

	for i := len(tr) - 1; i >= 0; i-- {
		if tr[i] == '#' {
			backspaces++
			tr[i] = ' '
			continue
		}

		if backspaces > 0 {
			tr[i] = ' '
			backspaces--
		}
	}

	i, j := 0, 0

	for i < len(sr) || j < len(tr) {

		for i < len(sr) && sr[i] == ' ' {
			i++
		}

		for j < len(tr) && tr[j] == ' ' {
			j++
		}

		if i >= len(sr) && j >= len(tr) {
			break
		}

		if (i >= len(sr) && j < len(tr)) || (j >= len(tr) && i < len(sr)) || sr[i] != tr[j] {
			return false
		}

		i++
		j++
	}

	return true
}
