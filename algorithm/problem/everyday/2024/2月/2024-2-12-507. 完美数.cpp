/*
507. 完美数
简单
相关标签
相关企业
对于一个 正整数，如果它和除了它自身以外的所有 正因子 之和相等，我们称它为 「完美数」。

给定一个 整数 n， 如果是完美数，返回 true；否则返回 false。



示例 1：

输入：num = 28
输出：true
解释：28 = 1 + 2 + 4 + 7 + 14
1, 2, 4, 7, 和 14 是 28 的所有正因子。
*/

class Solution {
public:
    bool checkPerfectNumber(int num) {
        int result = 0;
        int i = 1;
        while (i < num)
        {
            if (num % i == 0)
            {
                result += i;
            }
            i++;
        }
        if (result == num)
        {
            return true;
        }
        cout << result << endl;
        return false;
    }
};