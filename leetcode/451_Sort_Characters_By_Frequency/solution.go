package _451_Sort_Characters_By_Frequency

func frequencySort(s string) string {
	if len(s) < 3 {
		return s
	}

	freq, freqR := make(map[rune]uint), make(map[uint][]rune)

	for _, v := range s {
		freq[v]++
	}

	for i, v := range freq {
		if _, ok := freqR[v]; ok {
			freqR[v] = append(freqR[v], i)
		} else {
			freqR[v] = []rune{i}
		}
	}

	res := ""

	for i := uint(len(s)); i > 0; i-- {
		if bucket, ok := freqR[i]; ok {
			for x := 0; x < len(bucket); x++ {
				for j := uint(0); j < i; j++ {
					res += string(bucket[x])
				}
			}

		}
	}

	return res
}
