package other

import (
	"sort"
)

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。



示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
*/

// merge1 使用内置的sort包实现
func merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// merge2 双指针
func merge2(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0)
	// p1 record nums1; p2 record nums2
	var p1, p2 int
	for {
		if p1 == m {
			result = append(result, nums2[p2:]...)
			break
		}
		if p2 == n {
			result = append(result, nums1[p1:]...)
			break
		}
		if nums2[p2] > nums1[p1] {
			result = append(result, nums1[p1])
			p1++
		} else {
			result = append(result, nums2[p2])
			p2++
		}
	}
	copy(nums1[0:], result)
}

// merge3 3指针 降低内存使用
func merge3(nums1 []int, m int, nums2 []int, n int) {
	// p1指向nums1有效位置
	// p2指向nums2有效位置
	// p指向合成后的有效位置
	var p1, p2, p int = m - 1, n - 1, m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] >= nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	// 由于把数据处理到nums1 就算nums1还有元素 也不需要处理 本来就是有序
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p--
		p2--
	}
}
