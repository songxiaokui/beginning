--[[
1、字符串的表示方式
    - 单引号
    - 双引号
    - [[ 表示法
2、字符串长度获取
    - utf8.len 可以处理包含汉字的字符串长度
    - string.len 只能计算包含 ASCII 字符串编码的长度
--]]

local s1 = 'hello world'
local s2 = "hello 世界"
local s3 = [[
多行字符1
多行字符2
多行字符3
]]

print(utf8.len(s2))
print(string.len(s2))