/*
238. 除自身以外数组的乘积
中等
相关标签
相关企业
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。



示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]
*/

// go
func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    for i, _ := range result {
        result[i] = 1
    }
    for i:=1;i<n;i++ {
        result[i] = result[i-1]*nums[i-1]
    }
    temp := 1
    for i:=n-1;i>=0;i-- {
        result[i] = temp*result[i]
        temp = nums[i]*temp
    }
    return result
}

// 解题思路
// 使用结果数组初始化为1，第一次循环记录这个元素左侧的乘积 第二次反向循环 保存右侧积并缓存累积
class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        int n = nums.size();
        // 初始化结果数组
        vector<int> result(n, 1);
        // 保存当前元素左侧的乘积
        // i = 0 , v = 1
        // i = 1, v = result[i-1] * nums[i-1]
        for (int i = 1; i < n; i++)
        {
            result[i] = result[i-1] * nums[i-1];
        }
        // 逆向处理右侧乘积，并缓存当前乘积
        int temp(1);
        for (int i = n-1; i >= 0; i--)
        {
            result[i] = result[i] * temp;
            temp = temp * nums[i];
        }
        return result;
    }
};