/*
242. 有效的字母异位
相关标签
相关企业
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。



示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
*/
class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.size() != t.size())
        {
            return false;
        }
        vector<int> li(26, 0);
        for (int i = 0; i < s.size();i++)
        {
            li[s[i]-'a']++;
        }

        for (int i = 0; i < t.size();i++)
        {
            li[t[i]-'a']--;
        }
        for (int i = 0; i < 26; i++)
        {
            if (li[i] != 0)
            {
                return false;
            }
        }
        return true;
    }
};