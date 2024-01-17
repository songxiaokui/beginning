/*
2744. 最大字符串配对数目
已解答
简单
相关标签
相关企业
提示
给你一个下标从 0 开始的数组 words ，数组中包含 互不相同 的字符串。

如果字符串 words[i] 与字符串 words[j] 满足以下条件，我们称它们可以匹配：

字符串 words[i] 等于 words[j] 的反转字符串。
0 <= i < j < words.length
请你返回数组 words 中的 最大 匹配数目。

注意，每个字符串最多匹配一次。



示例 1：

输入：words = ["cd","ac","dc","ca","zz"]
输出：2
解释：在此示例中，我们可以通过以下方式匹配 2 对字符串：
- 我们将第 0 个字符串与第 2 个字符串匹配，因为 word[0] 的反转字符串是 "dc" 并且等于 words[2]。
- 我们将第 1 个字符串与第 3 个字符串匹配，因为 word[1] 的反转字符串是 "ca" 并且等于 words[3]。
可以证明最多匹配数目是 2 。
*/

class Solution {
public:
    int maximumNumberOfStringPairs(vector<string>& words) {
        int total = 0;
        map<string, int> mp;
        int length = words.size();
        for (int i = 0; i <= length-2;i++)
        {
            if (cache(words[i], mp))
            {
                continue;
            }
            for (int j = i+1; j <= length-1; j++)
            {
                if (compare(words[i], words[j]))
                {
                    total++;
                    mp[words[i]] = 0;
                    mp[words[j]] = 0;
                }
            }
        }
        return total;
    }

    bool cache(string s, map<string, int>& mp)
    {
        auto iter = mp.find(s);
        if (iter != mp.end())
        {
            return true;
        }
        else
        {
            return false;
        }
    }

    bool compare(string s1, string s2)
    {
        if (s1.size() != s2.size())
        {
            return false;
        }
        string s3 = s1 + s2;
        int left = 0;
        int right = s3.size()-1;
        while (left < right)
        {
            if (s3[left] != s3[right])
            {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
};