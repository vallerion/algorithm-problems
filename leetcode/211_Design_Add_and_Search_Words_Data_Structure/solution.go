package _211_Design_Add_and_Search_Words_Data_Structure

type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

// Constructor /** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{&TrieNode{make(map[rune]*TrieNode), false}}
}

func (this *WordDictionary) AddWord(word string) {
	this.Insert([]rune(word))
}

func (this *WordDictionary) Insert(word []rune) {
	root := this.root

	for i := 0; i < len(word); i++ {
		node, exists := root.children[word[i]]
		if exists == false {
			node = &TrieNode{make(map[rune]*TrieNode), false}
			root.children[word[i]] = node
		}
		root = node
	}

	root.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	root := this.root
	runeWord := []rune(word)
	queue := make([]*TrieNode, 0)
	queue = append(queue, root)
	var node *TrieNode

	i := 0 // index in runeWord

	for len(queue) > 0 {
		n := len(queue)

		if i >= len(runeWord) {
			return false
		}

		if runeWord[i] == '.' {
			for j := 0; j < n; j++ {
				for _, nn := range queue[j].children {
					queue = append(queue, nn)

					if i == len(runeWord)-1 && nn.isWord == true {
						return true
					}
				}
			}
		} else {
			for j := 0; j < n; j++ {
				exists := false
				node, exists = queue[j].children[runeWord[i]]

				if exists == true {
					queue = append(queue, node)

					if i == len(runeWord)-1 && node.isWord == true {
						return true
					}
				}
			}
		}

		queue = queue[n:]
		i++
	}

	if i == len(runeWord)-1 && node != nil && node.isWord == true {
		return true
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
