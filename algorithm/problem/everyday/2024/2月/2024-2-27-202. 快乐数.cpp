/*
202. 快乐数
相关标签
相关企业
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

示例 1：

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
*/

class Solution {
public:
    bool isHappy(int n) {
        map<int, bool> mp;
        while (n != 1)
        {
            if (find(mp, n))
            {
                break;
            }
            mp[n] = true;
            n = next(n);
        }
        return n == 1;
    }

    bool find(map<int, bool>& mp, int k)
    {
        auto iter = mp.find(k);
        if (iter != mp.end())
        {
            return true;
        }
        else
        {
            return false;
        }
    }

    int next(int n)
    {
        int nwn = 0;
        while (n != 0)
        {
            int a = n % 10;
            nwn += a * a;
            n = n / 10;
        }
        return nwn;
    }
};

// go
func isHappy(n int) bool {
    mp := make(map[int]struct{}, 0)
    for n != 1 {
        if _, has := mp[n]; has {
            goto END
        }
        mp[n] = struct{}{}
        n = next(n)
    }
    END:
    return n == 1
}


func next(n int) int {
    var result = 0
    for n != 0 {
        result += (n%10) * (n%10)
        n /= 10
    }
    return result
}