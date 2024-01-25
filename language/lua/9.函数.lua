--[[
函数是一等公民

定义:
    可选作用域[local] function 函数名(行参列表1, 行参2, ...)
        函数体
        return 返回值可以选
    end

函数作为参数传递

函数可以多返回值 如果使用变量接受超过返回值的变量 则值为 nil 字符串

函数的变长参数: ... 实现
    可以使用 {...}去接受参数 转换为表格进行处理
    也可以通过 select("#", ...) 获取可变参数的个数
    select(n, ...) 返回从起点 n 到结束的参数列表 如果将其进行赋值 则返回的是左侧的第一个元素

--]]

-- 函数定义
local function max(a, b)
    if a > b then
        return a
    else
        return b
    end
end

local a = 1
local b = 2
print(max(a, b))

-- 函数作为函数的参数
local function myprint(data)
    print("结果为: ", data)
end

local function add(a, b, func)
    local data = a + b
    func(data)
end

add(a, b, myprint)

-- 函数多返回值
-- string.find 返回字符串的开始索引和结束索引 所以 d3 == "nil"
d1, d2, d3 = string.find("sxkaust", "aust")
print(d1, d2, d3)

function sayHobby(...)
    -- 处理传入的参数
    local result = 0
    for k,v in pairs{...} do
        if type(v) == "function" then
            v()
        else
            result = result + v
        end
    end
    -- 计算平均值
    local avge = result/#{...}
    return result, avge
end

data, avge = sayHobby(1, 2, 3, function()print("hello world") end, 5)
print("总和为: ", data)
print("平均为: ", avge)

-- 使用 select 可以获取可变参数的个数
function getNumber(a, ...)
    return select("#", ...)
end

function getParams(n, ...)
    return select(n, ...)

end

print("参数个数为: ", getNumber(1, 2, 3, 4, 5))
print("从 2 开始的参数列表为: ", getParams(1, 2, 3, 4, 5))

do
    function foo(...)
        local n = select("#", ...)
        print("n = ", n)
        for i=1,n do
            print(select(i, ...))  -- 直接 print 是打印所有元素 预计底层存在一个迭代器进行数据输出
            local arg = select(i, ...) -- 将 select(i, ...) 赋值给变量 是取的左侧第一个元素
            print("arg: ", arg)
        end
    end
    foo(1, 2, 3, 4, 5)
end