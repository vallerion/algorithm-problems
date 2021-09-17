package _677_Map_Sum_Pairs

type MapSum struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	mapSum   map[string]int
}

func (tn *TrieNode) calcSum() int {
	total := 0

	for _, v := range tn.mapSum {
		total += v
	}

	return total
}

func (tn *TrieNode) pushMapSum(key string, value int) {
	tn.mapSum[key] = value
}

// Constructor /** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{&TrieNode{make(map[rune]*TrieNode), make(map[string]int)}}
}

func (this *MapSum) Insert(key string, val int) {
	root := this.root
	runeKey := []rune(key)

	for i := 0; i < len(runeKey); i++ {
		next, exists := root.children[runeKey[i]]
		if exists == false {
			next = &TrieNode{make(map[rune]*TrieNode), make(map[string]int)}
			root.children[runeKey[i]] = next
		}

		next.pushMapSum(key, val)
		root = next
	}
}

func (this *MapSum) Sum(prefix string) int {
	root := this.root
	runePrefix := []rune(prefix)

	for i := 0; i < len(runePrefix); i++ {
		next, exists := root.children[runePrefix[i]]
		if exists == false {
			return 0
		}

		root = next
	}

	return root.calcSum()
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
