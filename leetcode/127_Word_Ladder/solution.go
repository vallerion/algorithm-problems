package _127_Word_Ladder

// "a" -> a,b,c
// "c" -> a,b,c
// ["a","b","c"]
// a -> c

// level = 2
// "hot" -> hot,dot | level=3
// "dog" -> dog,dot | level=4
// "hot" -> hot,dot -> dog | return level-1 -> 3
// "dog" -> dog,dot

// ["hot","dog","dot"]
// hot -> dot -> dog

// "red" -> ted,rex -> tex,red -> tax
// "tax" -> tad -> ted
// ["ted","tex","red","tax","tad","den","rex","pee"]

func ladderLength(beginWord string, endWord string, wordList []string) int {
	check := existsInList(endWord, wordList)
	if check == false {
		return 0
	}

	level := 2

	beginMap, endMap := make(map[string]bool), make(map[string]bool)
	for i := 0; i < len(wordList); i++ {
		beginMap[wordList[i]] = true
		endMap[wordList[i]] = true
	}

	beginQueue, endQueue := make([]string, 0), make([]string, 0)
	beginQueue = append(beginQueue, beginWord)
	endQueue = append(endQueue, endWord)

	for len(beginQueue) > 0 && len(endQueue) > 0 {
		beginLen, endLen := len(beginQueue), len(endQueue)

		if isIntersecting(&beginMap, &endMap) {
			return level - 1
		}

		level++
		for i := 0; i < beginLen; i++ {
			for word, usable := range beginMap {
				if usable == true && couldUse(beginQueue[i], word) {
					if word == endWord {
						return level
					}

					beginQueue = append(beginQueue, word)
					beginMap[word] = false
				}
			}
		}

		level++
		for i := 0; i < endLen; i++ {
			for word, usable := range endMap {
				if usable == true && couldUse(endQueue[i], word) {
					if word == beginWord {
						return level
					}
					endQueue = append(endQueue, word)
					endMap[word] = false
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
