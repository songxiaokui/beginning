/*
1. 两数之和
已解答
简单
相关标签
相关企业
提示
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
*/

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int, int> mp;
        vector<int> result;
        for (int i = 0; i < nums.size(); i++)
        {
            if (mp.count(target-nums[i]))
            {
                result.push_back(mp[target-nums[i]]);
                result.push_back(i);
            }
            else
            {
                mp[nums[i]] = i;
            }
        }
        return result;
    }
};

// go
/*
func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)
    result := make([]int, 0)
    for i, v := range nums {
        if has, ok := mp[target-v]; ok {
            result = append(result, has, i)
            return result
        } else {
            mp[v] = i
        }
    }
    return result
}
*/

// python3
/*
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        mp = dict()
        result = list()
        for key, value in enumerate(nums):
            if mp.get(target-value) != None:
                result.append(mp[target-value])
                result.append(key)
                return result
            else:
                mp[value] = key
        return result
*/