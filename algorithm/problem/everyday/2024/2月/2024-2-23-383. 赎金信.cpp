/*
383. 赎金信
简单
相关标签
相关企业
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。
示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false
示例 2：

输入：ransomNote = "aa", magazine = "ab"
输出：false
示例 3：

输入：ransomNote = "aa", magazine = "aab"
输出：true
*/

// 解法1: 暴力法 统计ransomNote每个字符的次数 再统计magazine每个字符的次数 如果magazine存在ransomNote中的每个字符并且次数大于等于 则可构成
func canConstruct(ransomNote string, magazine string) bool {
    mm := make(map[rune]int)
    rm := make(map[rune]int)
    for _, v := range magazine {
        if _,has := mm[v]; !has {
            mm[v] = 1
        } else {
            mm[v] = mm[v]+1
        }
    }

    for _, v := range ransomNote {
        if _,has := rm[v]; !has {
            rm[v] = 1
        } else {
            rm[v] = rm[v]+1
        }
    }

    for k,v := range rm {
        data, has := mm[k]
        if !has {
            return false
        }
        if v > data {
            return false
        }
    }
    return true
}

// 解法2: 使用一个数组完成统计
// Go
func canConstruct(ransomNote string, magazine string) bool {
    if len(magazine) < len(ransomNote) {
        return false
    }
    li := make([]int, 26)
    for _, v := range magazine {
        li[byte(v)-'a']++
    }

    for _, v := range ransomNote {
        li[byte(v)-'a']--
        if li[byte(v)-'a'] < 0 {
            return false
        }
    }
    return true
}

// CPP
class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        if (magazine.size() < ransomNote.size())
        {
            return false;
        }
        vector<int> li(26);
        for (auto a: magazine)
        {
            li[a-'a']++;
        }

        for (auto a: ransomNote)
        {
            li[a-'a']--;
            if (li[a-'a'] < 0)
            {
                return false;
            }
        }
        return true;
    }
};