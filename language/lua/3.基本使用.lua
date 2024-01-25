#!/usr/local/bin/lua

-- 单行注释
-- lua -i 进入交换命令行

-- 多行注释
--[[
    lua 多行注释
--]]

-- 标识符
-- 标识符和常用语言一样 使用数字 字母 下划线构成 不嫩用数字开头
name = "sxk"
age = 18
high = 161.1
-- 格式化输出使用 string.format
message = string.format("姓名: %s, 年龄: %d, 身高: %.2f\n", name, age, high)
print(message)

-- 默认情况下 变量是全局变量
-- 删除一个变量 使用 nil
name = nil
print(name)
