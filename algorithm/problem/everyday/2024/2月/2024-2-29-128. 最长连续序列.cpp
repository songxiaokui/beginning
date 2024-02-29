/*
128. 最长连续序列
中等
相关标签
相关企业
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
*/

// 暴力法
class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        std::sort(nums.begin(), nums.end());
        int n = removeDuplicates(nums, nums.size());
        if (n <= 1)
        {
            return n;
        }
        int left = 0, right = 1;
        int result = 1;
        while (right < n)
        {
            if (nums[right]-nums[left] == right-left)
            {
                result = max(result, right-left+1);
                right++;
                continue;
            }
            // 更新left = right
            left = right;
            // right++
            right++;

        }

        return result;
    }

    int max(int a, int b) {
        return a > b ?a:b;
    }

    int removeDuplicates(vector<int>& arr, int n) {
    if (n == 0 || n == 1) {
        return n;
    }

    int j = 0; // j 记录去重后的数组长度

    // 从第二个元素开始，与前一个元素比较，如果不相同，则加入结果数组
    for (int i = 1; i < n; i++) {
        if (arr[i] != arr[j]) {
            arr[++j] = arr[i];
        }
    }

    // 返回去重后的数组长度
    return j + 1;
    }
};