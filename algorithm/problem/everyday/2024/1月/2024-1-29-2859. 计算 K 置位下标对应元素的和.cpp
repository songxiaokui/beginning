/*
2859. 计算 K 置位下标对应元素的和
已解答
简单
相关标签
相关企业
提示
给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

请你用整数形式返回 nums 中的特定元素之 和 ，这些特定元素满足：其对应下标的二进制表示中恰存在 k 个置位。

整数的二进制表示中的 1 就是这个整数的 置位 。

例如，21 的二进制表示为 10101 ，其中有 3 个置位。



示例 1：

输入：nums = [5,10,1,5,2], k = 1
输出：13
解释：下标的二进制表示是：
0 = 0002
1 = 0012
2 = 0102
3 = 0112
4 = 1002
下标 1、2 和 4 在其二进制表示中都存在 k = 1 个置位。
因此，答案为 nums[1] + nums[2] + nums[4] = 13 。
*/

// 暴力法
// 将索引转换为二进制字符串
// 计算 1 的数量是否== k
// 如果是 则加该索引对应的元素
// 否则 跳过

class Solution {
public:
    int sumIndicesWithKSetBits(vector<int>& nums, int k) {
        int result = 0;
        for (int i = 0; i < nums.size(); i++)
        {
            int a = nums[i];
            int count = convertNumber2Binary1Sum(i);
            if (isValid(count, k))
            {
                result += a;
            }
        }
        return result;
    }

     int convertNumber2Binary1Sum(int num)
    {
        int total = 0;
        while (num != 0)
        {
            int m = num % 2;
            if (m == 1)
            {
                total++;
            }
            num = num / 2;
        }
        return total;
    }

    bool isValid(int nums, int k)
    {
        if (nums == k)
        {
            return true;
        }
        return false;
    }
};

// go
/*
func sumIndicesWithKSetBits(nums []int, k int) int {
    var total int
    for i, v := range nums {
        if k == convert(i) {
            total += v
        }
    }
    return total
}

func convert(i int) int {
    var sum int
    for (i != 0) {
        if i % 2 == 1 {
            sum++
        }
        i /= 2
    }
    return sum
}
*/