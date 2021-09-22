package _648_Replace_Words

import "strings"

type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{&TrieNode{make(map[rune]*TrieNode), false}}
}

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func (trie *Trie) insert(word string) {
	runeWord := []rune(word)
	root := trie.root

	for i := 0; i < len(runeWord); i++ {
		node, exists := root.children[runeWord[i]]
		if exists == false {
			node = &TrieNode{make(map[rune]*TrieNode), false}
			root.children[runeWord[i]] = node
		}
		root = node
	}

	root.isWord = true
}

func (trie *Trie) search(word []rune) (bool, string) {
	root := trie.root

	for i := 0; i < len(word); i++ {
		node, exists := root.children[word[i]]
		if exists == false {
			return false, ""
		}
		if node.isWord == true {
			return true, string(word[:i+1])
		}
		root = node
	}

	return false, ""
}

func replaceWords(dictionary []string, sentence string) string {
	trie := newTrie()

	for i := 0; i < len(dictionary); i++ {
		trie.insert(dictionary[i])
	}

	runeSentence := []rune(sentence)
	wordStart := 0
	result := ""

	for i := 0; i < len(runeSentence); i++ {
		if runeSentence[i] == ' ' || i == len(runeSentence)-1 {

			currWord := runeSentence[wordStart:i]
			if i == len(runeSentence)-1 {
				// если слово последнее то индекс надо взять +1
				// иначе слайс возьмется как [wordStart:i)
				currWord = runeSentence[wordStart:i+1]
			}

			isFound, replaceWord := trie.search(currWord)

			if isFound == true {
				result += replaceWord + " "
			} else {
				result += string(currWord) + " "
			}

			wordStart = i + 1
		}
	}

	return strings.TrimRight(result, " ")
}
