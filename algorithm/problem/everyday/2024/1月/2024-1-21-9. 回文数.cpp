/*
9. 回文数
简单
相关标签
相关企业
提示
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。


示例 1：

输入：x = 121
输出：true
*/

class Solution {
public:
    bool isPalindrome(int x) {
        string s = to_string(x);
        int left = 0;
        int right = s.size()-1;
        while (left < right)
        {
            if (s[left]-s[right])
            {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
};

// go
func isPalindrome(x int) bool {
    s := strconv.Itoa(x)
    var left, right = 0, len(s)-1
    for left < right {
        if s[left] != s[right]{
            return false
        }
        left++
        right--
    }
    return true
}

// python3
class Solution:
    def isPalindrome(self, x: int) -> bool:
        s = str(x)
        left, right = 0, len(s)-1
        while left < right:
            if s[left] != s[right]:
                return False
            else:
                left += 1
                right-= 1
        return True