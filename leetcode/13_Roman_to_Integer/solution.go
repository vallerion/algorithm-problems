package _13_Roman_to_Integer

// MCMXCIV
// 1994

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// IX

func romanToInt(s string) int {
	sr := []rune(s)
	hp := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := hp[sr[len(sr)-1]]

	for i := len(sr) - 2; i >= 0; i-- {
		if hp[sr[i]] >= hp[sr[i+1]] {
			res += hp[sr[i]]
		} else {
			res -= hp[sr[i]]
		}
	}

	return res
}
