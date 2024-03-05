package _3月

import "strings"

/*
151. 反转字符串中的单词
中等
相关标签
相关企业
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。



示例 1：

输入：s = "the sky is blue"
输出："blue is sky the"
示例 2：

输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。
*/

func reverseWords(s string) string {
	sb := strings.Builder{}
	if len(s) == 0 || len(strings.ReplaceAll(s, " ", "")) == 0 {
		return ""
	}
	// 双指针 结束条件1. 遇到空格 2. 索引为0
	n := len(s) - 1
	move, flag := n, n+1
	for move >= 0 {
		if move == 0 || string(s[move]) == " " {
			if string(s[move:flag]) == " " {
				flag = move
				move--
				continue
			}
			left := move + 1
			if move == 0 {
				left = move
			}
			if move == 0 && string(s[move]) == " " {
				left = move + 1
			}
			sb.WriteString(s[left:flag])
			sb.WriteString(" ")
			flag = move
			move--
			continue
		}
		move--
	}
	return strings.TrimSuffix(sb.String(), " ")
}
