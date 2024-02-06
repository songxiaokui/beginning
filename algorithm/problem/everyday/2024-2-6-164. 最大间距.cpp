/*
164. 最大间距
中等
相关标签
相关企业
给定一个无序的数组 nums，返回 数组在排序之后，相邻元素之间最大的差值 。如果数组元素个数小于 2，则返回 0 。

您必须编写一个在「线性时间」内运行并使用「线性额外空间」的算法。



示例 1:

输入: nums = [3,6,9,1]
输出: 3
解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
示例 2:

输入: nums = [10]
输出: 0
解释: 数组元素个数小于 2，因此返回 0。
*/

class Solution {
public:
    int maximumGap(vector<int>& nums) {
        int result = 0;
        // 冒泡排序会超时
        // bobble(nums, nums.size());

        // 使用自带的排序库
        std::sort(nums.begin(), nums.end());
        // 处理数据
        for (int i = 0; i+1 < nums.size(); i++ )
        {
            int tmp = nums[i+1] - nums[i];
            result = result > tmp ? result:tmp;
        }

        return result;
    }

    // bobble sort
    void bobble(vector<int>& nums, int n)
    {
        for (int i = 0; i < n-1 ; i++)
        {
            for (int j=i; j < n; j++)
            {
                if (nums[i] > nums[j])
                {
                    int temp = nums[j];
                    nums[j] = nums[i];
                    nums[i] = temp;
                }
            }
        }
    }
};

// go
func maximumGap(nums []int) int {
    sort.Ints(nums)
    result := 0
    for i:=0;i+1<len(nums);i++ {
        if nums[i+1] - nums[i] > result {
            result = nums[i+1] - nums[i]
        }
    }
    return result
}