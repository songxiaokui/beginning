/*
4. 寻找两个正序数组的中位数
困难
相关标签
相关企业
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/

// go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l := len(nums1)+len(nums2)
    mid1 := 0
    mid2 := 0
    if l % 2 == 0 {
        mid1 = l/2 - 1
        mid2 = l/2
    } else {
        mid1 =  l/2
        mid2 = l/2
    }

    result := make([]int, len(nums1)+len(nums2))
    a := 0 // point nums1
    b := 0 // point nums2
    i := 0 // mid value index

    for a < len(nums1) && b < len(nums2) {
        if nums1[a] < nums2[b] {
            result[i] = nums1[a]
            a++
        } else {
            result[i] = nums2[b]
            b++
        }

        // 判断i是否到达
        if i == mid2 {
            if mid1 == mid2 {
                return  float64(result[i])
            }
            return (float64(result[i]) + float64(result[i-1]))/2

        }
        i++
    }

    // 处理没达到的元素
    if a == len(nums1) {
        // a 全部处理完毕
        for b < len(nums2) {
            result[i] = nums2[b]
            b++
            if i == mid2 {
            if mid1 == mid2 {
                return  float64(result[i])
                }
                return (float64(result[i]) + float64(result[i-1]))/2
            }
        i++
        }
    }

    if b == len(nums2) {
        // b 全部处理完毕
        for a < len(nums1) {
            result[i] = nums1[a]
            a++
            if i == mid2 {
            if mid1 == mid2 {
                return  float64(result[i])
            }
            return (float64(result[i]) + float64(result[i-1]))/2
        }
        i++
        }
    }
    return 0
}

// 解题思路: 边比较边判断是否是中间值

// C++
class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int n1l = nums1.size();
        int n2l = nums2.size();
        int l = n1l + n2l;

        // midstore
        vector<int> result(n1l+n2l, 0);

        // handle index
        int m1 = 0;
        int m2 = 0;

        // target index
        if ((n1l+n2l)%2)
        {
            // 奇数
            m2 = l / 2;
            m1 = l / 2;
        } else
        {
            // 偶数
            m2 = l / 2;
            m1 = l / 2 - 1;
        }

        int n1 = 0;
        int n2 = 0;
        int i = 0;

        while (n1 < n1l && n2 < n2l)
        {
            if (nums1[n1] < nums2[n2])
            {
                result[i] = nums1[n1];
                n1++;
            } else
            {
                result[i] = nums2[n2];
                n2++;
            }
            // 判断是否处理到中间值
            if (i == m2)
            {
                if (m1 == m2)
                {
                    return (float) result[i];
                }
                return ((float) result[i] + (float) result[i-1]) / 2;
            }
            i++;
        }

        // 当元素处理为空
        if (n1 == n1l)
        {
            // nums2 还有元素
            while (n2 < n2l)
            {
                result[i] = nums2[n2];
                if (i == m2)
            {
                if (m1 == m2)
                {
                    return (float) result[i];
                }
                return ((float) result[i] + (float) result[i-1]) / 2;
            }
                i++;
                n2++;
            }
        }

        if (n2 == n2l)
        {
            // nums1 还有元素
            while (n1 < n1l)
            {
                result[i] = nums1[n1];
                if (i == m2)
            {
                if (m1 == m2)
                {
                    return (float) result[i];
                }
                return ((float) result[i] + (float) result[i-1]) / 2;
            }
                i++;
                n1++;
            }
        }
        return 0;
    }
};

// C++ 简化代码 配合goto
class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int n1l = nums1.size();
        int n2l = nums2.size();
        int l = n1l + n2l;

        // midstore
        vector<int> result((n1l+n2l)/2+1, 0);

        // handle index
        int m1 = 0;
        int m2 = 0;

        // target index
        if ((n1l+n2l)%2)
        {
            // 奇数
            m2 = l / 2;
            m1 = l / 2;
        } else
        {
            // 偶数
            m2 = l / 2;
            m1 = l / 2 - 1;
        }

        int n1 = 0;
        int n2 = 0;
        int i = 0;

        while (n1 < n1l && n2 < n2l)
        {
            if (nums1[n1] < nums2[n2])
            {
                result[i++] = nums1[n1++];

            } else
            {
                result[i++] = nums2[n2++];

            }
            if (i-1 == m2)
            {
                i--;
                goto END;
            }
        }

        // 当元素处理为空
        if (n1 == n1l)
        {
            // nums2 还有元素
            while (n2 < n2l)
            {
                result[i++] = nums2[n2++];
                if (i-1 == m2)
                {
                    i--;
                    goto END;
                }
            }
        }

        if (n2 == n2l)
        {
            // nums1 还有元素
            while (n1 < n1l)
            {
                result[i++] = nums1[n1++];
                if (i-1 == m2)
                {
                    i--;
                    goto END;
                }
            }
        }

        END:
        if (m1 == m2)
        {
            return (float) result[i];
        }
        return ((float) result[i] + (float) result[i-1]) / 2;
    }
};