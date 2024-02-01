--[[
metatable 元表
表无法实现两个表的操作
元表可以改变元表的行为

1. 建立元表 setmetatable(t1, {__index={}})
2. 元表的查询 __index(如果表存在 则直接返回 否则查询是否存在__index设置元表 如果存在 则从元表查询如此往复)
3. 元表的更新 __newindex(如果表存在key则更新 否则调用__newindex元表更新)
4. 元表的基本操作 + - * / % - concat equal lt gt
5. call
--]]

-- 1. 建立元表 setmetatable({},{})
t1 = { name = "sxk", age = 18 }
t2 = { password = "123" }
t3 = setmetatable(t1, { __index = t2 })
-- t3 = setmetatable(t1, t2) -- 不设置__index 无法在t2中查找
print(t3.name)
print(t3.age)
print(t3.password)
-- 如果在t1中查找不到元素 并且设置了t1的元表并且存在__index 则会在元表中查找 否则返回nil
-- 实质如下
t4 = setmetatable(t1, { __index = function(t1, key)
    if key == "password" then
        return "123456"
    end
    return nil
end })

print(t4.password)
-- 相当于多了一个查询路径 __index = {password="123456"}

-- 更新元表
t5 = { name = "s", age = 1 }
t6 = {}
t8 = {}
t7 = setmetatable(t5, { __newindex = t6 })
t7.name = "sxk"
print(t7.name)
t7.password = "124"
t8.password = "sss"
print(t7.password) -- 无法设置 因为存在元表设置 __newindex 对已经存在的key进行更新 则不会调用元表的更新
print(t6.password) -- t6为元表 则会将数据赋值到元表中
print(t8.password)
-- 使用rawset 进行初始表更新

t9 = setmetatable(t5, { __newindex = function(t5, key, value)
    rawset(t5, key, "--" .. value .. "--")
end })
t9.name = "000"
t9.on = "on"
print(t9.name) -- 存在的key 不调用__newindex逻辑
print(t9.on) -- --on--

-- 加
t10 = { 1, 2, 4 }
t11 = { 3, 9, 11 }
t10 = setmetatable(t10, { __add = function(tb1, tb2)
    -- 第一个参数为+前一个数据 第二个为后一个数据
    l1 = #tb1
    for i = 1, #tb2 do
        table.insert(tb1, tb2[i]) -- 默认追加到最后
    end
    return tb1
end })

t10 = t10 + t11
for k, v in pairs(t10) do
    print("k=" .. k .. "->value=" .. v)
end

-- __call 重载的是表的()
t12 = setmetatable(t10, { __call = function(tb1)
    -- 第一个参数为+前一个数据 第二个为后一个数据
    local result = 0
    for _, v in pairs(tb1) do
        result = result + v
    end
    return result
end })
print(t12())