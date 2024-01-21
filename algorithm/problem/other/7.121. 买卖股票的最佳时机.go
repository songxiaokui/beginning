package other

import "math"

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
示例 1：

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
*/

func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	var slow, fast, result = 0, 1, 0
	for fast < length {
		d := prices[fast] - prices[slow]
		if d < 0 {
			slow = fast
			fast++
			continue
		}
		if d > result {
			result = d
		}
		fast++
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

func maxProfit2(prices []int) int {
	var mi, ma = math.MaxInt, 0
	for _, v := range prices {
		ma = max(v-mi, ma)
		mi = min(v, mi)
	}
	return ma
}
