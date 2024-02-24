/*
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。



示例1:

输入: pattern = "abba", s = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", s = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false
*/

// Go
func wordPattern(pattern string, s string) bool {
    s1 := strings.Split(s, " ")
    if len(pattern) != len(s1) {
        return false
    }
    pm := make(map[byte]string, 0)
    sm := make(map[string]byte, 0)
    if len(s1) != len(pattern) {
        return false
    }
    for i:=0;i<len(s1);i++ {
        a,b := pattern[i],s1[i]
        if v, has := pm[a]; has && v != b {
            return false
        }
        if v1, has2 := sm[b]; has2 && v1 != a {
            return false
        }
        pm[a] = b
        sm[b] = a
    }
    return true
}

// C++
class Solution {
public:
    bool wordPattern(string pattern, string s) {
        vector<string>* l = m_split(s, ' ');
        int n = l->size();
        if (n != pattern.size())
        {
            return false;
        }
        map<char, string> pm;
        map<string, char> sm;
        for (int i = 0; i < n; i++)
        {
            char c = pattern[i];
            string s = (*l)[i];
            if (pm.count(c)>0 && pm[c] != s || sm.count(s)>0&&sm[s] != c)
            {
                return false;
            }
            pm[c] = s;
            sm[s] = c;
        }
        delete l;
        l = nullptr;
        return true;
    }

    vector<string>* m_split(string s, char tk)
    {
        vector<string>* l = new vector<string>();
        std::istringstream iss(s);
        string token;
        while (std::getline(iss, token, tk))
        {
            l->push_back(token);
        }
        return l;
    }
};