/*
66. 加一
简单
相关标签
相关企业
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]
*/

// 主要考虑最后一个数字+1后的进位问题
class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        int carry = 0;
        for (int i = digits.size()-1; i >= 0; i--)
        {
            if (i == digits.size()-1)
            {
                digits[i]++;
            }
            int d = carry + digits[i];
            int r = d % 10;
            carry = d / 10;
            digits[i] = r;
        }
        if (carry != 0)
        {
           // 插入首部
           digits.insert(digits.begin(), carry);
        }
        return digits;
    }
};

// go
func plusOne(digits []int) []int {
    length := len(digits)
    result := make([]int, 0)
    carry := 0
    for i:=length-1;i>=0;i--{
        if i == length-1 {
            digits[i]++
        }
        // 处理值
        num := carry + digits[i]
        r := num % 10
        // 处理进位
        carry = num / 10
        digits[i] = r
    }
    if carry != 0 {
        result = append(result, carry)
        result = append(result, digits...)
    } else {
        return digits
    }
    return result
}