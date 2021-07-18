package main

// a a b a a b a a a a b  a  a  b  а  a  a  b
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17
// 0 0 0 0 0 0 0 0 0 0 0  0  0  0  0  0  0  0
// _ _

// _ _ -> i=1, j=i-1
// 0 1 0 0 0 0 0 0 0 0 0  0  0  0  0  0  0  0

// i=2,j=1 (j=p[2-1]=p[1]=1)

func main() {
	//ыва
}

func prefix(str string) []int {
	s := []rune(str)
	prefixes := make([]int, len(str))

	for i := 1; i < len(str); i++ {
		j := prefixes[i-1]

		if s[i] == s[j] {
			j++
		}

		prefixes[i] = j
	}

	return prefixes
}
