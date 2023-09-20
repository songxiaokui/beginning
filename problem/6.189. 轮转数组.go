package problem

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
*/

// 利用slice特性实现
func rotate1(nums []int, k int) {
	length := len(nums)
	var realK = 1
	if k != 1 {
		realK = k % length
	}
	result := make([]int, length)
	result = append(nums[length-realK:], nums[:length-realK]...)
	copy(nums, result)
}

// 该方案O(n^2) 超时
func rotate2(nums []int, k int) {
	length := len(nums)
	var realK = 1
	if k != 1 {
		realK = k % length
	}
	var index int
	var right = length - 1
	for index < realK {
		temp := nums[right]
		saveRight := right
		for right > 0 {
			nums[right] = nums[right-1]
			right--
		}
		nums[0] = temp
		index++
		right = saveRight
	}
}

// 使用三次翻转实现
func rotate3(nums []int, k int) {
	length := len(nums)
	realK := k % length
	revers(nums, 0, length-1)
	revers(nums, 0, realK-1)
	revers(nums, realK, length-1)
}

func revers(array []int, left, right int) {
	for left < right {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
}
