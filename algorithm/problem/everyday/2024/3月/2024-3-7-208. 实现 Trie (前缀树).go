package _3月

type Trie struct {
	child map[rune]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{
		child: make(map[rune]*Trie, 0),
		isEnd: false,
	}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, e := range word {
		if _, has := node.child[e]; !has {
			node.child[e] = &Trie{child: make(map[rune]*Trie), isEnd: false}
		}
		node = node.child[e]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, e := range word {
		if _, has := node.child[e]; !has {
			return false
		}
		node = node.child[e]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, e := range prefix {
		if _, has := node.child[e]; !has {
			return false
		}
		node = node.child[e]
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
