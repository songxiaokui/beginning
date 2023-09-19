package problem

/*
26
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
*/

// 使用map进行保存是否出现过
func removeDuplicates(nums []int) int {
	hm := make(map[int]byte)
	var index int
	for _, v := range nums {
		if _, has := hm[v]; !has {
			nums[index] = v
			index++
			hm[v] = '1'
		}
	}
	return index
}

// 使用快慢指针
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}
	var slow, fast = 1, 1
	for ; fast < n; fast++ {
		// or nums[fast] != nums[slow-1]
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
