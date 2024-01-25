--[[
变量
    1.变量使用前 先定义 否则直接使用为 nil
    2.Lua 变量默认为全局变量
    3.建议多使用局部变量

变量类型:
    1. 全局变量
        全局可见
    2. 局部变量
       local 关键词修饰 从申明到结束都可见
       无法在作用域外访问函数内部的局部变量
    3. 表中的域

table两种取值方式:
    1. 索引取值
    2. 点取值

table 对象是引用传递
--]]

a = 10 -- 全局变量
local b = 10 -- 局部变量

function call()
    print("函数内部访问全局变量 a= " .. a)
    local c = 100
    print("函数内部访问局部变量 b= " .. b)
    b = 200
    print("修改局部变量 b = " .. b)
end

call()
print("局部变量 b = " .. b)

-- 无法访问函数内部定义的局部变量 c
print(c)

do
    local x, y = 10
    print(x, y)  -- 10 nil
    -- 多赋值 从左依次赋值 如果不够 默认为 nil
end

-- 实现交换
function swap(a1, b1)
    print("内部交换前: ".."a1="..a1.." b1="..b1)
    a1, b1 = b1, a1
    print("内部交换后: ".."a1="..a1.." b1="..b1)
    return a1, b1
end

a1, b1 = 10, 99
print("交换前: ".."a1="..a1.." b1="..b1)
a1, b1 = swap(a1, b1)
print("交换后: ".."a1="..a1.." b1="..b1)

-- table的两种取值方式
table1 = {1, 99, 100, name="dd", age=19, 6}
print(table1[1])
print(table1.name)

-- 注意 table 默认是引用传递
function swapArrayElem(arr, left, right)
    arr[left], arr[right] = arr[right], arr[left]
end
print(table1[1], table1[4])
swapArrayElem(table1, 1, 4)
print(table1[1], table1[4])