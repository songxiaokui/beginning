/*
169. 多数元素
简单
相关标签
相关企业
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。



示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2
*/

func majorityElement(nums []int) int {
    if len(nums) <=1 {
        return nums[0]
    }
    count := len(nums)/2
    mp := make(map[int]int, 0)
    for _, v := range nums {
        if c, has := mp[v]; !has {
            mp[v]=1
        } else {
            if c+1>count {
                return v
            } else {
                mp[v]++
            }
        }
    }
    return 0
}

// 库方法: 先排序 返回中位数
func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}

// Cpp
class Solution {
public:
    int majorityElement(vector<int>& nums) {
        std::sort(nums.begin(), nums.end());
        return nums[nums.size()/2];

    }
};

class Solution {
public:
    int majorityElement(vector<int>& nums) {
        if (nums.size() == 1)
        {
            return nums[0];
        }
        int count = nums.size()/2;
        map<int, int> mp;
        for (auto i: nums)
        {
            if (++mp[i]>count)
            {
                return i;
            }
        }
        return 0;
    }
};