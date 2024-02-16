/*
670. 最大交换
中等
相关标签
相关企业
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

示例 1 :

输入: 2736
输出: 7236
解释: 交换数字2和数字7。
示例 2 :

输入: 9973
输出: 9973
解释: 不需要交换。
注意:

给定数字的范围是 [0, 108]
*/

/*
解题思路:
    1. 暴力法
    把所有的交换情况例举出来 比较最大值 然后返回
    每次交换后记得把数据还原回原来的状态
*/

// C++
class Solution {
public:
    int maximumSwap(int num) {
        string s = std::to_string(num);
        int n = s.size();
        int maxNum = num;
        for (int i = 0; i < n; i++)
        {
            for ( int j = i+1; j <n;j++)
            {
                swap(s, i, j);
                maxNum = max(maxNum, stoi(s));
                swap(s, i, j);
            }
        }
        return maxNum;
    }
    void swap(string &a, int b, int c)
    {
        char temp = a[b];
        a[b] = a[c];
        a[c] = temp;
    }

    template<class T>
    T max(T a, T b)
    {
        return a > b ? a: b;
    }
};

// go
/*
func maximumSwap(num int) int {
    var s []rune = []rune(strconv.Itoa(num))
    maxNumber := num
    for i, _ := range s {
        for j:=i+1; j <len(s); j++ {
            swap(s, i, j)
            n, _ := strconv.Atoi(string(s))
            maxNumber = max(maxNumber, n)
            swap(s, i, j)
        }
    }
    return maxNumber
}

func swap(s []rune, a, b int) {
    s[a], s[b] = s[b], s[a]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
*/