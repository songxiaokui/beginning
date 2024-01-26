--[[
1、字符串的表示方式
    - 单引号
    - 双引号
    - [[ 表示法
2、字符串长度获取
    - utf8.len 可以处理包含汉字的字符串长度
    - string.len 只能计算包含 ASCII 字符串编码的长度
3、string 包常用操作
    upper 将字符串转大写
    lower 将字符串转小写
    gsub 字符串指定替换
    find 查找子字符串 返回开始到结束的索引
    reverse 字符串反转
    format 格式化输出
    char 将整数转字符并连接
    byte 将字符串转数字 只转第一个
    len 获取字符串的长度
    rep 字符串的复制
    .. 字符串连接功能
    gmatch 字符串匹配 返回的是匹配的迭代器 如果不存在 则为 nil
    match 字符串匹配 只匹配第一个
    sub 实现字符串的截取 sub(字符串, 开始位置, [, 结束位置)
4、字符串的正则匹配
    符合正常的正则匹配规则
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

-- 字符串的常见操作
-- string.upper 将字符串转大写
local s4 = string.upper(s1)
print(s4)

-- string.lower 将字符串转小写
print(string.lower(s4))

-- string.gsub 在字符串中替换
print(string.gsub("abcabc", "a", "x", 1))

--find 查找子字符串 返回开始到结束的索引
print(string.find("sxk hello !", "he"))

--reverse 字符串反转
print(string.reverse("hello world"))

--format 格式化输出
print(string.format("%d, %.2f", 10, 9.9111))

--char 将整数转字符并连接
print(string.char(72, 73, 77))

--byte 将字符串转数字 只转第一个
print(string.byte("hello world"))
print(string.byte("h"))

--len 获取字符串的长度
print(string.len("hello"))

--rep 字符串的复制
print(string.rep("sxk", 10))

--... 字符串连接功能
print("sxk " .. "is good!")

--gmatch 字符串匹配 返回的是匹配的迭代器 如果不存在 则为 nil
for wold in string.gmatch("sxk is good", "%a+") do
    print(wold)
end

--match 字符串匹配 只匹配第一个
print(string.match("sxk is good!", "%a+"))

--sub 实现字符串的截取 sub(字符串, 开始位置, [, 结束位置)
-- 从指定位置截取到结尾
local s5 = "hello world"
print(string.sub(s5, 3))
-- 从指定位置结束到指定位置
print(string.sub(s5, 1, 2))
-- 截取倒数5个
print(string.sub(s5, -5))
-- 索引越界 截取全部
print(string.sub(s5, -100))