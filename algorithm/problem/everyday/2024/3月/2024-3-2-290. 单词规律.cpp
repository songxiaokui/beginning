/*
290. 单词规律
简单
相关标签
相关企业
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。



示例1:

输入: pattern = "abba", s = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", s = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false
*/

// go
func wordPattern(pattern string, s string) bool {
    pm := make(map[byte]string, 0)
    sm := make(map[string]byte, 0)
    sl := strings.Split(s, " ")
    if len(pattern) != len(sl) {
        return false
    }
    for i := range pattern {
        a, b := pattern[i], sl[i]
        if (pm[a]!="" && pm[a] != b) || (sm[b]> 0&&sm[b] != a) {
            return false
        }
        pm[a] = b
        sm[b] = a
    }
    return true
}

