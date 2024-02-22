/*
189. 轮转数组
中等
相关标签
相关企业
提示
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。



示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
*/

// 超时
class Solution1 {
public:
    void rotate(vector<int>& nums, int k) {
        // 如果k是nums.size()的倍数 则等于原数组
        int n = nums.size();
        if (k % n == 0)
        {
            return;
        }
        // 创建一个移动缓存数组 然后移动 前面的数据 然后移动缓存中的数据即可
        int nk = k % n;
        while (nk > 0)
        {
            // 搬运数据的同时 移动前面的数据
            // 存储最后一个数据
            int data = nums[n-1];
            int c = nums.size()-1;
            // 移动数据
            while (c > 0)
            {
                nums[c] = nums[c-1];
                c--;
            }
            // 赋值在第一个位置
            nums[c] = data;
            nk--;
        }
    }
};

// 解法一：超时，主要在于移动数据次数随着k的变大而变大

class Solution {
public:
    void rotate(vector<int>& nums, int k) {
        // 如果k是nums.size()的倍数 则等于原数组
        int n = nums.size();
        if (k % n == 0)
        {
            return;
        }
        // 创建一个移动缓存数组 然后移动 前面的数据 然后移动缓存中的数据即可
        int nk = k % n;
        vector<int> cache;
        for (int i = n-nk; i < n; i++)
        {
            // cache-> [5, 6, 7]
            cache.push_back(nums[i]);
        }

        // 进行两次移动元素
        // 1. 移动原数组
        for (int i = n-1; i-nk>=0; i--)
        {
            nums[i] = nums[i-nk];
        }
        // 2. 移动缓存数组
        for (int i = 0; i < cache.size();i++)
        {
            nums[i] = cache[i];
        }
    }
};

// 解法二: 使用一个缓存数组先保存被移动的数据，然后先移动原数组 然后进行缓存数组的赋值