package _208_Implement_Trie

type Trie struct {
	root *node
}

type node struct {
	children  map[rune]*node
	isEndWord bool
}

// Constructor /** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{&node{make(map[rune]*node), false}}
}

// Insert /** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this.root
	runeWord := []rune(word)

	for i := 0; i < len(runeWord); i++ {
		next, exists := root.children[runeWord[i]]
		if exists == false {
			next = &node{make(map[rune]*node), false}
			root.children[runeWord[i]] = next
		}

		root = next
	}

	root.isEndWord = true
}

// Search /** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this.root
	runeWord := []rune(word)

	for i := 0; i < len(runeWord); i++ {
		next, exists := root.children[runeWord[i]]
		if exists == false || (i == len(runeWord)-1 && next.isEndWord == false) {
			return false
		}
		root = next
	}

	return true
}

// StartsWith /** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this.root
	runeWord := []rune(prefix)

	for i := 0; i < len(runeWord); i++ {
		next, exists := root.children[runeWord[i]]
		if exists == false {
			return false
		}
		root = next
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
