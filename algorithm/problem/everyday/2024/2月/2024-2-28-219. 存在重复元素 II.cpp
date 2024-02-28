/*
219. 存在重复元素 II
简单
相关标签
相关企业
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
示例 1：

输入：nums = [1,2,3,1], k = 3
输出：true
示例 2：

输入：nums = [1,0,1,1], k = 1
输出：true
示例 3：

输入：nums = [1,2,3,1,2,3], k = 2
输出：false
*/

// 暴力法
// 两次循环 判断值和索引
// 结果: 超出时间限制
class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        int n = nums.size();
        if (n <= 1)
        {
            return false;
        }
        for (int i = 0; i < n-1; i++)
        {
            for (int j = i+1; j < n; j++)
            {
                if (nums[i]==nums[j] && j-i<=k)
                {
                    return true;
                }
            }
        }
        return false;
    }
};

// O(n) 一次循环+map缓存数据
class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        int n = nums.size();
        if (n <= 1)
        {
            return false;
        }
        map<int, int> mp;
        for (int i = 0; i < n; i++)
        {
            if (mp.count(nums[i]) == 0)
            {
                mp[nums[i]] = i;
                continue;
            }
            int index = mp[nums[i]];
            if (i - index <= k)
            {
                return true;
            }
            else
            {
                mp[nums[i]] = i;
            }
        }
        return false;
    }
};