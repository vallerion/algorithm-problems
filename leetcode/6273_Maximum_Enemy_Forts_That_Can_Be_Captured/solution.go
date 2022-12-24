package _6273_Maximum_Enemy_Forts_That_Can_Be_Captured

// [1,0,0,-1,0,0,0,0,1]

func captureForts(forts []int) int {
	maxCaptured := 0

	for i := 0; i < len(forts); i++ {
		if forts[i] == 0 {
			continue
		}
		for j := i + 1; j < len(forts); j++ {
			if forts[i] == forts[j] {
				break
			}
			if forts[i] != forts[j] && forts[j] != 0 {
				if j-i-1 > maxCaptured {
					maxCaptured = j - i - 1
				}
			}
		}
	}

	return maxCaptured
}
