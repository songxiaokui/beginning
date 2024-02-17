/*
205. 同构字符串
简单
相关标签
相关企业
给定两个字符串 s 和 t ，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。



示例 1:

输入：s = "egg", t = "add"
输出：true
示例 2：

输入：s = "foo", t = "bar"
输出：false
示例 3：

输入：s = "paper", t = "title"
输出：true
*/

// go
func isIsomorphic(s string, t string) bool {
    sm := make(map[byte]byte, 0)
    tm := make(map[byte]byte, 0)
    for i := range s {
        a, b := s[i], t[i]
        if sm[a]>0 && sm[a] != b || tm[b]>0 && tm[b] != a {
            return false
        }
        sm[a] = b
        tm[b] = a
    }
    return true
}

// C++
class Solution {
public:
    bool isIsomorphic(string s, string t) {
        map<char, char> sm;
        map<char, char> tm;
        for (int i = 0; i < s.length(); i++)
        {
            char a = s[i];
            char b = t[i];
            if (find_element(sm, a) && sm[a] != b || find_element(tm, b) && tm[b] != a)
            {
                return false;
            }
            sm[a] = b;
            tm[b] = a;
        }
        return true;
    }

    bool find_element(map<char, char>& mp, char elem)
    {
        if (mp.count(elem) > 0)
        {
            return true;
        }
        return false;

    }
};