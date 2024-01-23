/*
20. 有效的括号
已解答
简单
相关标签
相关企业
提示
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：

输入：s = "()"
输出：true
*/

// go
/*
func isValid(s string) bool {
    mp := make(map[rune]rune)
    mp[')'] = '('
    mp['}'] = '{'
    mp[']'] = '['
    li := make([]rune, 0)
    for _, v := range []rune(s) {
        data, has := mp[v]
        if !has {
            li = append(li, v)
            continue
        }
        // 如果遇到右括号
        if len(li) == 0 || li[len(li)-1] != data {
            return false
        } else {
            // 匹配上了，则删除该末尾元素
            li = li[0:len(li)-1]
        }
    }
    if len(li) != 0 {
        return false
    }
    return true
}

// 解题思路
// 将所有的左符号压入栈中  如果遇到右括号 则弹出匹配的左括号 如果第一个不是 则 false 如果最后栈不为空 也不是
*/

class Solution {
public:
    bool isValid(string s) {
        map<char, char> mp;
        mp[')'] = '(';
        mp['}'] = '{';
        mp[']'] = '[';

        stack<char> st;
        for (auto& a: s)
        {
            cout << a << endl;
            if (mp.count(a) == 0)
            {
                st.push(a);
                continue;
            }
            // 如果存在
            if (st.empty() || st.top() != mp[a])
            {
                return false;
            }
            // 匹配上，弹出栈顶元素
            st.pop();
        }

        if (st.empty())
        {
            return true;
        }
        return false;
    }
};