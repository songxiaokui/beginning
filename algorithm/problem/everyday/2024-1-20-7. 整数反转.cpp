/*
7. 整数反转
已解答
中等
相关标签
相关企业
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。


示例 1：

输入：x = 123
输出：321
*/

// 解题思路
// 取模后 * 10 + 当前模长

// 使用字符串包会报错
class Solution1 {
public:
    int reverse(int x) {
        int sign = 1;
        if (x<0)
        {
            sign = -1;
        }
        string s = to_string(x);
        rev(s, s.size()-1);
        return sign* stoi(s);
    }

    void rev(string& s, int l)
    {
        int left = 0;
        int right = l;
        while (left < right)
        {
            char temp = s[left];
            s[left] = s[right];
            s[right] = temp;
            left++;
            right--;
        }
    }
};

// 正确解法
class Solution {
public:
    int reverse(int x) {
        int d = 0;
        while (x != 0)
        {
            if (d > INT_MAX / 10 || d < INT_MIN / 10) {
                return 0;
            }
            d = d*10 + x % 10;
            x = x / 10;
        }
        return d;
    }
};

//go
func reverse(x int) int {
    var d int = 0
    for (x != 0){
        if temp := int32(d); (temp*10)/10 != temp {
			return 0
		}
        d = (d * 10 + x % 10)
        x = x / 10
    }
    return int(d)
}