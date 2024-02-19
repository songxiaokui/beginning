/*
26. 删除有序数组中的重复项
已解答
简单
相关标签
相关企业
提示
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
判题标准:

系统会用下面的代码来测试你的题解:

int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = removeDuplicates(nums); // 调用

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
如果所有断言都通过，那么您的题解将被 通过。



示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

*/

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int hand_index = 0;
        int move_index = 0;
        if (!nums.size())
        {
            return 0;
        }
        map<int, bool> mp;
        while (move_index < nums.size())
        {
            if (mp.count(nums[move_index])) {
                move_index++;
                continue;
            }
            // 若不在
            mp[nums[move_index]] = true;
            nums[hand_index] = nums[move_index];
            hand_index++;
            move_index++;
        }
        return hand_index;
    }
};

//go
func removeDuplicates(nums []int) int {
    mp := make(map[int]struct{}, 0)
    var hand_index, move_index int = 0, 0
    for move_index < len(nums) {
        if _, ok := mp[nums[move_index]]; ok {
           move_index++
            continue
        }
        mp[nums[move_index]] = struct{}{}
        nums[hand_index] = nums[move_index]
        hand_index++
        move_index++

    }
    return hand_index
}

//python3
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        mp = dict()
        hand, move = 0, 0
        while move < len(nums):
            if mp.get(nums[move]) != None:
                move += 1
                continue
            else:
                mp[nums[move]] = 1
                nums[hand] = nums[move]
                hand += 1
                move += 1
        return hand

// 不使用map
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        if (nums.size() <= 1)
        {
            return nums.size();
        }
        int hand = 0;
        int move = 0;
        int val = -1000000;
        while (move < nums.size())
        {
            if (nums[move]-val==0)
            {
                move++;
                continue;
            }
            nums[hand++] = nums[move++];
            val=nums[hand-1];
        }
        return hand;
    }
};

// go
func removeDuplicates(nums []int) int {
    var hand, move, val int = 0, 0, -100000
    for move < len(nums) {
        if nums[move]-val == 0 {
            move++
            continue
        }
        nums[hand] = nums[move]
        hand++
        move++
        val = nums[hand-1]
    }
    return hand
}