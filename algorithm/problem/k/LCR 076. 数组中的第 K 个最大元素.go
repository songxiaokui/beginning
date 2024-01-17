package k

import (
	"math"
	"sort"
)

func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest(nums []int, k int) int {
	sl := make([]int, k)
	for i, _ := range sl {
		sl[i] = math.MinInt64
	}

	for _, v := range nums {
		insert(sl, v, k-1)
	}
	return sl[k-1]
}

func insert(el []int, ele int, k int) {
	if el[k] > ele {
		return
	}
	location := 0
	// 找到插入的位置
	for i := k; i >= 0; i-- {
		if el[i] > ele {
			location = i + 1
			break
		}
	}
	// 移动元素
	for i := k; i-1 >= location; i-- {
		el[i] = el[i-1]
	}
	// 插入数值
	el[location] = ele
}
