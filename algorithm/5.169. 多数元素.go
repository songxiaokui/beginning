package algorithm

import "sort"

/*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素
示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2
*/

func majorityElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[length-1]
	}
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
		if mp[v] > length/2 {
			return v
		}
	}
	return -1
}

func majorityElement2(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[length-1]
	}
	sort.Ints(nums)
	return nums[(length / 2)]
}
