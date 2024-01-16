/*
70. 爬楼梯
已解答
简单
相关标签
相关企业
提示
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？



示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/

// 暴力法
class Solution1 {
public:
    int climbStairs(int n) {
        vector<int> li = vector<int>{1,2};
        vector<int> count = vector<int>{0};
        int state = 0;
        update(state, n, count, li);
        li.clear();
        return count[0];
    }
    //穷举
    void update(int state, int n, vector<int>& count, vector<int>& li)
    {
        // 如果到达顶部
        if (state == n)
        {
            // 方案+1
            count[0]++;
            return;
        }
        // 如果没有到达顶部, 继续执行爬楼梯方案
        for (auto iter = li.begin(); iter != li.end(); iter++)
        {
            // 如果已经越界，就不在继续执行
            if (state + (*iter) > n)
            {
                continue;
            }
            // 递归查找
            update(state+(*iter), n, count, li);
        }

    }
};

// 递归法
class Solution2 {
public:
    int climbStairs(int n) {
        return dfs(n);
    }

    int dfs(int n)
    {
        if (n == 1 || n == 2)
        {
            return n;
        }

        return dfs(n-1) + dfs(n-2);
    }
};

// 迭代法(半动态规划)
class Solution3 {
public:
    int climbStairs(int n) {
        int n1 = 1;
        int n2 = 2;
        if (n <= 2)
        {
            return n;
        }
        for (int i = 3; i <= n; i++)
        {
            int temp = n1;
            n1 = n2;
            n2 += temp;
        }
        return n2;

    }

};

// 递归法+解决重叠子问题=>记忆搜索法(自顶向底)
// 将大问题拆解成小问题 直到拆解到已知最小问题的解
class Solution4 {
public:
    int climbStairs(int n) {
        map<int, int> mp;
        return dfs(n, mp);
    }

    int dfs(int n, map<int, int>& mp)
    {
        if (n == 1 || n == 2)
        {
            return n;
        }

        if (mp.find(n) != mp.end())
        {
            return mp.at(n);
        } else
        {
            int count = dfs(n-1, mp) + dfs(n-2, mp);
            mp[n] = count;
            return count;
        }

    }
};

// 动态规划(自底向上)
// 用已知问题的解 构建最后解决的问题
// 动态规划
class Solution {
public:
    int climbStairs(int n) {
        // 已知问题解
        if (n <= 2)
        {
            return n;
        }
        vector<int> dp(n+1);
        dp[1] = 1;
        dp[2] = 2;
        // 找到状态转移方程
        // dp[n] = dp[n-1] + dp[n-2]
        for (int i = 3; i <= n; i++)
        {
            dp[i] = dp[i-1] + dp[i-2];
        }
        return dp[n];
    }
};