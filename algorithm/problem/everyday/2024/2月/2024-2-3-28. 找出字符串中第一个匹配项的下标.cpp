/*
28. 找出字符串中第一个匹配项的下标
已解答
简单
相关标签
相关企业
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。



示例 1：

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
*/

// go
func strStr(haystack string, needle string) int {
    for index, _ := range haystack {
        if len(needle)+index > len(haystack) {
            return -1
        }
        if haystack[index] == needle[0] && haystack[index:len(needle)+index] == needle {
            return index
        }
    }
    return -1
}

//cpp
class Solution {
public:
    int strStr(string haystack, string needle) {
        int l = haystack.size();
        int start = 0;
        int l2 = needle.size();
        while (start < l)
        {
            if (start+l2 > l)
            {
                return -1;
            }
            if (check(haystack, needle, start))
            {
                return start;
            }
            start++;
        }
        return -1;
    }

    bool check(string a, string b, int c)
    {
        int i = 0;
        while (i < b.size())
        {
            if (a[c+i] != b[i])
            {
                return false;
            }
            i++;
        }
        return true;
    }
};

