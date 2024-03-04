package _3月

/*
14. 最长公共前缀
简单
相关标签
相关企业
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/

// 使用trie前缀树保存数据 然后获取前缀长度

type TrieNode struct {
	child  map[rune]*TrieNode
	isEnd  bool
	weight int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		child:  make(map[rune]*TrieNode, 0),
		isEnd:  false,
		weight: 0,
	}
}

func (tn *TrieNode) insert(s string) {
	node := tn
	for _, e := range s {
		if _, has := node.child[e]; !has {
			node.child[e] = NewTrieNode()
		}
		node = node.child[e]
		node.weight++
	}
	node.isEnd = true
}

func (tn *TrieNode) getPrefixIndex(c int) int {
	return recursion(tn, 0, c)
}

func recursion(root *TrieNode, index int, c int) int {
	if len(root.child) == 1 {
		n := &TrieNode{}
		for k, _ := range root.child {
			n = root.child[k]
		}
		if n.weight == c {
			return 1 + recursion(n, index+1, c)
		}
		return 0
	}
	if len(root.child) != 1 {
		return 0
	}
	return 0
}

func longestCommonPrefix(strs []string) string {
	trie := NewTrieNode()
	for _, v := range strs {
		if v == "" {
			return ""
		}
		trie.insert(v)
	}
	l := trie.getPrefixIndex(len(strs))
	if l == 0 {
		return ""
	}
	return strs[0][0:l]
}
