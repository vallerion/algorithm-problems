package main

import "fmt"

func main() {
	var lines int

	fmt.Scanf("%d\n", &lines)

	for i := 0; i < lines; i++ {
		var n, k, s int // 10 7 6
		fmt.Scanf("%d %d %d\n", &n, &k, &s)

		// ans1 - with restart
		// ans2 - without restart
		ans1, ans2 := k+n, k-1
		// 17, 6

		if s >= k {
			ans2 = s - 1
			ans2 += n - ans2 + 1
		} else {
			ans2 += k - s     // 6 + (7-6) = 7
			ans2 += n - s + 1 // 12
		}

		fmt.Printf("Case #%d: %d\n", i+1, min(ans1, ans2))
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
