/*
27. 移除元素
已解答
简单
相关标签
相关企业
提示
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。



说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
int len = removeElement(nums, val);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}


示例 1：

输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
*/
//C++
class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int hand = 0;
        int move = 0;
        while (move < nums.size())
        {
            if (nums[move]-val == 0)
            {
                move++;
                continue;
            }
            nums[hand] = nums[move];
            hand++;
            move++;
        }
        return hand;
    }
};

//go
func removeElement(nums []int, val int) int {
    var hand, move int = 0, 0
    for move < len(nums){
        if nums[move]-val == 0 {
            move++
            continue
        }
        nums[hand] = nums[move]
        hand++
        move++
    }
    return hand
}

//python3
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        hand, move = 0, 0
        while move < len(nums):
            if nums[move]-val==0:
                move += 1
            else:
                nums[hand] = nums[move]
                hand+=1
                move+=1
        return hand