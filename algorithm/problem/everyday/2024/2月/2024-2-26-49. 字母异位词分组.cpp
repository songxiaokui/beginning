/*
49. 字母异位词分组
中等
相关标签
相关企业
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]
*/

// 解法1: 暴力法
class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.size() != t.size()) {
            return false;
        }
        if (s.empty()) { // 处理空字符串的情况
            return true;
        }
        vector<int> li(26, 0);
        for (int i = 0; i < s.size(); i++) {
            li[s[i]-'a']++;
        }
        for (int i = 0; i < t.size(); i++) {
            li[t[i]-'a']--;
        }
        for (int i = 0; i < 26; i++) {
            if (li[i] != 0) {
                return false;
            }
        }
        return true;
    }

    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> result;
        map<string, bool> handle;
        for (int i = 0; i < strs.size(); i++) {
            if (handle.count(strs[i]) > 0) {
                continue;
            }
            vector<string> temp;
            temp.push_back(strs[i]);
            for (int j = i+1; j < strs.size(); j++) {
                if (isAnagram(strs[i], strs[j])) {
                    temp.push_back(strs[j]);
                    handle[strs[i]] = true;
                    handle[strs[j]] = true;
                }
            }
            if (temp.size() != 0) {
                result.push_back(temp);
            }
        }
        return result;
    }
};

// 最后会超时,两次循环对比是否是异位词 然后加到临时数组

// 解法二: 异位词排序后应该相等 一次循环 使用map保存中间结果 最后将map汇合到结果中
class Solution {
public:

    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> result;
        map<string, vector<string>> cache;
        for (auto s: strs)
        {
            string temp = s;
            sort(s.begin(), s.end());
            if (cache.count(s)>0)
            {
                cache[s].push_back(temp);
            } else
            {
                vector<string> l;
                l.push_back(temp);
                cache[s].push_back(temp);
            }
        }
        for (auto iter = cache.begin(); iter != cache.end();iter++)
        {
            result.push_back((*iter).second);
        }
        return result;
    }
};

// go 使用string 转rune 然后使用 sort.Slice进行排序
func groupAnagrams(strs []string) [][]string {
    result := make([][]string, 0)
    mp := make(map[string][]string, 0)
    for _, v := range strs {
        // 字符串排序
        data := v
        tl := []rune(v)
        sort.Slice(tl, func(i, j int) bool {
            return tl[i] < tl[j]
        })
        nvs := string(tl)
        mp[nvs] = append(mp[nvs], data)
    }

    for _, v := range mp {
        result = append(result, v)
    }
    return result
}