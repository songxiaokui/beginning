/*
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

实现词典类 WordDictionary ：

WordDictionary() 初始化词典对象
void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。


示例：

输入：
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
输出：
[null,null,null,null,false,true,true,true]

解释：
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // 返回 False
wordDictionary.search("bad"); // 返回 True
wordDictionary.search(".ad"); // 返回 True
wordDictionary.search("b.."); // 返回 True
*/

struct TrieNode {
    vector<TrieNode*> child;
    bool isEnd;
    TrieNode() {
        this->child = vector<TrieNode*>(26, nullptr);
        this->isEnd = false;
    }
};

void insert(string word, TrieNode* root)
{
    TrieNode* node = root;
    for (auto w: word)
    {
        if (node->child[w-'a'] == nullptr)
        {
            node->child[w-'a'] = new TrieNode();
        }
        node = node->child[w-'a'];
    }
    node->isEnd = true;
}

class WordDictionary {
private:
    TrieNode* root;
public:
    WordDictionary() {
        this->root = new TrieNode();
    }

    void addWord(string word) {
        insert(word, this->root);
    }

    bool search(string word) {
        return dfs(word, 0, this->root);
    }

    bool dfs(const string & word, int index, TrieNode* root)
    {
        if (index == word.size())
        {
            return root->isEnd;
        }
        char c = word[index];
        // 如果是字符a-z
        if (c >= 'a' && c <= 'z')
        {
            TrieNode* n = root->child[c-'a'];
            if (n != nullptr && dfs(word, index+1, n))
            {
                return true;
            }
        }
        // 如果是.
        else if (c == '.')
        {
            for (auto n1: root->child)
            {
                if (n1 != nullptr && dfs(word, index+1, n1))
                {
                    return true;
                }
            }
        }
        return false;
    }
};

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary* obj = new WordDictionary();
 * obj->addWord(word);
 * bool param_2 = obj->search(word);
 */

 //  Go
 /*
 type TrieNode struct {
     Child [26]*TrieNode
     IsEnd bool
 }

 func NewTrieNode() *TrieNode{
     return &TrieNode{
         Child: [26]*TrieNode{},
         IsEnd: false,
     }
 }

 type WordDictionary struct {
     root *TrieNode
 }


 func Constructor() WordDictionary {
     return WordDictionary{
         root: NewTrieNode(),
     }
 }


 func (this *WordDictionary) AddWord(word string)  {
     node := this.root
     for _, e := range word {
         if node.Child[byte(e)-'a'] == nil {
             node.Child[byte(e)-'a'] = NewTrieNode()
         }
         node = node.Child[byte(e)-'a']
     }
     node.IsEnd = true
 }


 func (this *WordDictionary) Search(word string) bool {
     return dfs(word, 0, this.root)
 }

 func dfs(word string, index int, node *TrieNode) bool {
     if len(word) == index {
         return node.IsEnd
     }
     c := byte(word[index])
     if c >= 'a' && c <= 'z' {
         n := node.Child[c-'a']
         if n != nil && dfs(word, index+1, n) {
             return true
         }
     } else if c == '.' {
         for _, f := range node.Child {
             if f != nil && dfs(word, index+1, f) {
                 return true
             }
         }
     }
     return false

 }


 /**
  * Your WordDictionary object will be instantiated and called as such:
  * obj := Constructor();
  * obj.AddWord(word);
  * param_2 := obj.Search(word);
  */
 */