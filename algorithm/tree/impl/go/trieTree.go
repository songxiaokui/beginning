package _go

import (
	"errors"
)

type Tire struct {
	root *TrieNode
}

func NewTrie() *Tire {
	return &Tire{
		root: nil,
	}
}

type TrieNode struct {
	child map[rune]*TrieNode
	isEnd bool
}

func NewTrieNode(sign rune, isEnd bool) *TrieNode {
	child := make(map[rune]*TrieNode)
	child[sign] = &TrieNode{
		child: child,
		isEnd: false,
	}
	return &TrieNode{
		child: child,
		isEnd: isEnd,
	}
}

// Insert add a string into current tree
func (t *Tire) Insert(s string) error {
	if t.Search(s) {
		return errors.New("insert element is exists")
	}

	node := t.root
	for _, e := range s {
		if t.root == nil {
			t.root = NewTrieNode(e, false)
			node = t.root
			continue
		}
		if _, has := node.child[e]; !has {
			node.child[e] = NewTrieNode(e, false)
		}
		node = node.child[e]
	}
	if node != nil {
		node.isEnd = true
	}
	return nil
}

// Delete remove string from tree
func (t *Tire) Delete(s string) error {
	if t.root == nil {
		return errors.New("root is nil")
	}
	if !t.Search(s) {
		return errors.New("element is not exists")
	}

	return nil
}

// Search find element from tree
func (t *Tire) Search(s string) bool {
	if t.root == nil {
		return false
	}
	node := t.root
	for _, e := range s {
		if _, has := node.child[e]; !has {
			return false
		}
		node = node.child[e]
	}
	return node.isEnd
}

// StartWithPrefix prefix is exists
func (t *Tire) StartWithPrefix(prefix string) bool {
	if t.root == nil {
		return false
	}
	node := t.root
	for _, e := range prefix {
		if _, has := node.child[e]; !has {
			return false
		}
	}
	return true
}
