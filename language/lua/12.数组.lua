--[[
数组的基本操作
1. 数组定义
    使用 table 存储类型相同的一组数据

--]]

-- 1.数组定义
local list1 = {100, 200 ,300, 400}

-- 遍历数组
for i=1, #list1 do
    print("值: ", list1[i])
end
-- 访问索引和值
for k,v in pairs(list1) do
    print("key: "..k.." value: ".. v)
end

-- 通过索引修改值
list1[1] = 10000
print(list1[1])

-- 追加元素
list1[#list1+1] = 999
print(list1[#list1])

-- 删除元素
table.remove(list1, 1)
for i=1, #list1 do
    print("值: ", list1[i])
end

-- 多维数组
list2 = {}
for i=1, 10 do
    list2[i] = {}
    for j=1, 2 do
       list2[i][j] = j
    end
end

for i=1,#list2 do
    for j=1,#list2[i] do
        print(list2[i][j])
    end
end