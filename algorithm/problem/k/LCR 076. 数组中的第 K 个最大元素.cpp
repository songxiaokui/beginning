
/*
LCR 076. 数组中的第 K 个最大元素
已解答
中等
相关标签
相关企业
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。



示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
*/

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k)
    {
        vector<int> el(k, INT_MIN);
        for (int i = 0; i < nums.size(); i++)
        {
            insert(el, nums[i], k-1);
        }
        return el[k - 1];
    }

    // 插入数据
    void insert(std::vector<int>& el, int ele, int k)
    {
        // 最后一个比当前大 直接去掉
        if (el[k] >= ele)
        {
            return;
        }

        int localtion = 0;

        for (int i = k; i >= 0; i--)
        {
            if (el[i] > ele)
            {
                localtion = i + 1;
                break;
            }
        }

        for (int i = k; i-1 >= localtion; i--)
        {
            el[i] = el[i - 1];
        }

        el[localtion] = ele;
    }
};

// 使用C++的内置 sort 包实现
class Solution1 {
public:
    int findKthLargest(vector<int>& nums, int k) {
        std::sort(nums.begin(), nums.end(), std::greater<int>());
        return nums[k-1];
    }
};

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        std::sort(nums.begin(), nums.end());
        return nums[nums.size()-k];
    }
};