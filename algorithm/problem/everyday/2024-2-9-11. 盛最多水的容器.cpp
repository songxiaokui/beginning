/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。
*/

class Solution {
public:
    int maxArea(vector<int>& height) {
        int left = 0, right = height.size()-1;
        int result = 0;
        while (left < right)
        {
            int area = (right - left) * min(height[left], height[right]); // 计算面积
            result = max(result, area);
            if (height[left] < height[right])
            {
                left++;
            }
            else
            {
                right--;
            }
        }
        return result;
    }
    int max(int a, int b)
    {
        return a > b ? a : b;
    }
    int min(int a, int b)
    {
        return a < b ? a : b;
    }
};

// go
/*
func maxArea(height []int) int {
    left, right, result := 0, len(height)-1, 0
    for left < right {
        area := (right-left )* min(height[left], height[right])
        result = max(result, area)
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
*/