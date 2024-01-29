/*
228. 汇总区间
相关标签
相关企业
给定一个  无重复元素 的 有序 整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b
示例 1：

输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
*/

class Solution {
public:
    vector<string> summaryRanges(vector<int>& nums) {
        vector<string> result;
        int i = 0, l = nums.size();

        while (i < l)
        {
            int record = i;
            while (i+1 < l && nums[i] + 1 == nums[i+1])
            {
                i++;
            }
            if (record==i)
            {
                // 只有一个元素
                result.push_back(std::to_string(nums[record]));
            } else {
                // 从start->i
                result.push_back(std::to_string(nums[record])+"->"+std::to_string(nums[i]));
            }
            i++;
        }
        return result;
    }
};

// go
func summaryRanges(nums []int) []string {
    result := make([]string, 0)
    n := len(nums)
    i := 0
    for (i < n) {
        start:=i
        for i+1<n && nums[i]+1==nums[i+1] {
            i++;
        }
        if start == i {
            // 只有一个元素
            result = append(result, strconv.Itoa(nums[start]))
        } else {
            // 开始指向结束
            result = append(result, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[i]))
        }
        i++
    }
    return result
}