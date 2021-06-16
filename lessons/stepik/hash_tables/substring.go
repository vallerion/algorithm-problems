//package hash_tables
package main

import (
	"fmt"
)

/*
 * Task #5.3
 * @see ./statements.pdf
 */

const q = 10e+7
const mm = 1000
const xx = 1 // todo: сделать чтобы работало при x=263

// h(S) = S[i] * x^i + S[i+1] * x^i + S[i+2] * x^i
func hashPattern(s []rune) int {
	h := 0

	for i := 0; i < len(s); i++ {
		h = h + int(s[i])*pow2(xx, i)
		//h = h + int(s[i])*int(math.Pow(float64(xx), float64(i)))
	}

	return h % mm
}

func pow2(a, b int) int {
	if b == 0 || a == 1 {
		return 1
	}
	if b == 1 {
		return a
	}

	res := a

	for b > 1 {
		res = (res % q) * (a % q)
		b--
	}

	return res % q
}

func main() {
	var patternInput, textInput string

	fmt.Scanf("%s\n", &patternInput)
	fmt.Scanf("%s\n", &textInput)

	pattern := []rune(patternInput)
	text := []rune(textInput)

	hPattern := hashPattern(pattern)
	h := 0
	matches := make([]int, 0)
	hashes := make(map[int]int, 0)

	for i := 0; i < len(text); i++ {
		power := i
		if i >= len(pattern)-1 {
			power = len(pattern) - 1
		}

		hashes[i] = int(text[i]) * pow2(xx, power)
		//hashes[i] = (int(text[i]) % q) * (int(math.Pow(float64(xx), float64(i))) % q)
		h = h + hashes[i]

		if i >= len(pattern)-1 {
			//fmt.Println(power, h%mm, hPattern, int(text[i-len(pattern)+1]), int(text[i]), h-int(text[i-len(pattern)+1]))

			if h%mm == hPattern && matchCharacters(pattern, text, i-len(pattern)+1, i) {
				matches = append(matches, i-len(pattern)+1)
			}

			//h = h - int(text[i-len(pattern)+1])*int(math.Pow(float64(xx), float64(i-len(pattern)+1)))
			//h = (h - int(text[i-len(pattern)+1])) / xx
			//h = (h - hashes[i-len(pattern)+1]) / xx
			h = h - hashes[i-len(pattern)+1]
			//h = h - int(text[i-len(pattern)+1]) - pow2(xx, len(pattern))
		}
	}

	//fmt.Println(hashes)

	for i := 0; i < len(matches); i++ {
		fmt.Print(matches[i])
		fmt.Print(" ")
	}
}

func matchCharacters(pattern, text []rune, start, end int) bool {
	match := true
	pi := 0

	for start <= end {
		if text[start] != pattern[pi] {
			match = false
			break
		}
		start++
		pi++
	}

	return match
}
