/*
2719. 统计整数数目
困难
相关标签
相关企业
提示
给你两个数字字符串 num1 和 num2 ，以及两个整数 max_sum 和 min_sum 。如果一个整数 x 满足以下条件，我们称它是一个好整数：

num1 <= x <= num2
min_sum <= digit_sum(x) <= max_sum.
请你返回好整数的数目。答案可能很大，请返回答案对 109 + 7 取余后的结果。

注意，digit_sum(x) 表示 x 各位数字之和。



示例 1：

输入：num1 = "1", num2 = "12", min_num = 1, max_num = 8
输出：11
解释：总共有 11 个整数的数位和在 1 到 8 之间，分别是 1,2,3,4,5,6,7,8,10,11 和 12 。所以我们返回 11 。
示例 2：

输入：num1 = "1", num2 = "5", min_num = 1, max_num = 5
输出：5
解释：数位和在 1 到 5 之间的 5 个整数分别为 1,2,3,4 和 5 。所以我们返回 5 。
*/

/*
* 解法1: 暴力穷举法
* 解题思路:
    1. 将num1 和 num2 转换为整数
    2. 一个一个遍历 然后计算这个值每个位数之和
    3. 判断计算后的数是否满足条件min_sum <= digit_sum(x) <= max_sum.
    4. 满足+1
    5. 返回总数 % (int)(pow(10, 9) + 7)
* 结果:
    整数转换越界
*/
// 暴力解法1
class Solution1 {
public:
    int count(string num1, string num2, int min_sum, int max_sum) {
        int n1 = s2i(num1);
        int n2 = s2i(num2);
        int total = 0;
        while (n1 <= n2)
        {
            int temp = n1;
            temp = sum(temp);
            if (temp <= max_sum && temp >= min_sum)
            {
                total++;
            }
            n1++;
        }
        return total % (int)(pow(10, 9)+7);
    }

    // 将字符串转换为整数
    int s2i(string s)
    {
        return std::stoi(s);
    }

    // 计算数字每个位数之和
    int sum(int a)
    {
        int d = 0;
        while (a)
        {
            int temp = a;
            a /= 10;
            d += temp % 10;
        }
        return d;
    }
};