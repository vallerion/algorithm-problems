package _127_Word_Ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	check := existsInList(endWord, wordList)
	if check == false {
		return 0
	}

	hp := make(map[string]bool)
	for i := 0; i < len(wordList); i++ {
		hp[wordList[i]] = false
	}

	beginQueue, endQueue := make([]string, 0), make([]string, 0)
	beginQueue = append(beginQueue, beginWord)
	endQueue = append(endQueue, endWord)

	level := 1

	for len(beginQueue) > 0 && len(endQueue) > 0 {
		beginLen, endLen := len(beginQueue), len(endQueue)
		level++

		//if isIntersecting(&beginMap, &endMap) {
		//	return level
		//}

		for i := 0; i < beginLen; i++ {
			for word, usable := range hp {
				if usable == false && couldUse(beginQueue[i], word) {
					return level
				} else if couldUse(beginQueue[i], word) {
					beginQueue = append(beginQueue, word)
					hp[word] = false
				}
			}
		}

		for i := 0; i < endLen; i++ {
			for word, usable := range hp {
				if usable == false && couldUse(endQueue[i], word) {
					return level
				} else if couldUse(endQueue[i], word) {
					endQueue = append(endQueue, word)
					hp[word] = false
				}
			}
		}

		beginQueue = beginQueue[beginLen:]
		endQueue = endQueue[endLen:]
	}

	return 0
}

func isIntersecting(beginMap, endMap *map[string]bool) bool {
	for word, usable := range *beginMap {
		if v, ok := (*endMap)[word]; ok && v == false && usable == false {
			return true
		}
	}

	return false
}

// ------- BFS --------

func ladderLengthBFS(beginWord string, endWord string, wordList []string) int {
	check := existsInList(endWord, wordList)
	if check == false {
		return 0
	}

	mp := make(map[string]bool)
	for i := 0; i < len(wordList); i++ {
		mp[wordList[i]] = true
	}

	queue := make([]string, 0)
	queue = append(queue, beginWord)

	level := 0

	for len(queue) > 0 {
		n := len(queue)

		for i := 0; i < n; i++ {
			if queue[i] == endWord {
				return level + 1
			}

			for word, usable := range mp {
				if usable == false {
					continue
				}

				if couldUse(queue[i], word) {
					queue = append(queue, word)
					mp[word] = false
				}
			}
		}

		queue = queue[n:]
		level++
	}

	return 0
}

func couldUse(word, listsWord string) bool {
	counter := 0

	for i := 0; i < len(word); i++ {
		if word[i] != listsWord[i] {
			counter++
		}
		if counter > 1 {
			return false
		}
	}

	return true
}

func existsInList(word string, list []string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == word {
			return true
		}
	}

	return false
}
