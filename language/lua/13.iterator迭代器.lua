--[[
迭代器
    用来遍历模板库容器的元素
实现:
    需要保存三个值 迭代函数 状态常量 控制变量
    迭代函数会返回2个值 一个是结束状态nil 或者 返回k，v给变量
    pairs 就是一个迭代函数 内置的泛型迭代器迭代函数

迭代函数:
    有状态
    无状态
--]]

-- 无状态迭代函数 只使用 状态常量和控制变量 进行处理下一个元素
function square(maxCount, currentNumber)
    if currentNumber < maxCount then
        currentNumber = currentNumber + 1
        return currentNumber, currentNumber^2
    end
end

for k, v in square, 3, 0 do
    print(k .." -> ".. v)
end

-- 模拟pairs 实现有状态迭代器 需要使用闭包进行状态维护 多状态
function mypairs(tb)
    local l = #tb
    local init = 0
    return function()
        if init < l then
            init = init + 1
            return init, tb[init]
        end
    end
end

local l1 = {100, 200 ,300 ,400}
for k,v in mypairs(l1) do
    print(k.."->".. v)
end

