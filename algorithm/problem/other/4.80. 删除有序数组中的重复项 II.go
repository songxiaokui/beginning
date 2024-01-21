package other

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。
*/

func removeDuplicatesII1(nums []int) int {
	mp := make(map[int]int)
	var index int
	for _, v := range nums {
		data, has := mp[v]
		if !has || data != 2 {
			nums[index] = v
			mp[v]++
			index++
		}
	}
	return index
}

func removeDuplicatesII2(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	var slow, fast = 2, 2
	for ; fast < length; fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
