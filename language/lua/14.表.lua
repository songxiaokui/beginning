--[[
表(table)的基本使用:
1. 定义与初始化
    使用构造器进行初始化 {}
    tb = {}
2. table 对象赋值给一个对象，两个对象共享内存，一个对象修改，另一个对象访问的值也会发生变化
3. 一个变量清空表 不会影响另外一个变量
4. table可以使用迭代器遍历 for k,v in pairs(tb) do 执行体 end

表的常用操作:
1. 连接 concat
2. 插入与删除 insert remove
3. 排序 sort
4. 获取最大值
5. 获取长度
--]]

-- 表构造器初始化
mytable = {}
mytable[1] = "Lua"
mytable["m"] = "hello"

newtable = mytable

print(mytable[1])
print(newtable[1])

newtable[1] = "go"
print(mytable[1])
print(newtable[1])

newtable = nil
print(newtable)
print(mytable["m"])

-- table可以使用迭代器遍历
for k, v in pairs(mytable) do
    print(k..v)
end

--表连接
a = {10, 20, 30}
print(table.concat(a))  -- 默认以空格连接
print(table.concat(a, "+")) -- 指定字符连接
print(table.concat(a, "-", 1, 2)) -- 指定字符连接并 指定连接的范围

-- 插入
a[#a+1] = 100
print(#a)
print(a[#a])

table.insert(a, 9900) -- 默认在末尾插入
print(a[#a])

table.insert(a, 1, 888) -- 指定索引插入,更新
print(#a)
print("---"..a[1])

--删除
table.remove(a, #a)
print(a[#a])

-- 排序
print(table.concat(a, ","))
table.sort(a)
print(table.concat(a, ","))

