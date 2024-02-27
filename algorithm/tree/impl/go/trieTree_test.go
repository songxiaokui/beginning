package _go

import "testing"

func TestTrieTree(t *testing.T) {
	trie := NewTrie()
	t.Log("搜索字符串: sxk: ", trie.Search("sxk"))
	t.Log("插入字符串: sxk: ", trie.Insert("sxk"))
	t.Log("插入字符串: sxi: ", trie.Insert("sxi"))
	t.Log("重复插入字符串: sxi: ", trie.Insert("sxi"))
	t.Log("搜索字符串: sxk: ", trie.Search("sxk"))
	t.Log("搜索字符串: sxi: ", trie.Search("sxi"))
	t.Log("搜索前缀字符串: sx: ", trie.StartWithPrefix("sx"))
	t.Log("搜索前缀字符串: sxdd: ", trie.StartWithPrefix("sxdd"))
	t.Log("搜索前缀字符串: s: ", trie.StartWithPrefix("s"))
	t.Log("删除sxk: sxk: ", trie.Delete("sxk"))
	t.Log("搜索字符串: sxi: ", trie.Search("sxi"))
	t.Log("搜索字符串: sxk: ", trie.Search("sxk"))
}
