--[[
循环方式:
    1. while 条件为真执行循环体
    2. for 循环执行
    3. repeat...until 重复执行直到为真

goto:
    代码跳转
    定义 label方式 ::标签名称::
    使用: goto 标签名称
--]]

-- while 循环
n = 100
total = 0
while (n > 0) do
    if (n % 2 == 0) then
        total = total + n
    end
    n = n - 1
end
print("0-100 偶数和为: ", total)

-- for 循环
-- for 初始条件,结束条件 do {循环执行体} end
for n=1, 100 do
    if (n%2==0) then
        total = total - n
    end
end

print("total 最后的值为 0 is ", 0 == total)

-- repeat {重复执行循环体} until (终止条件)
p = 1
repeat
    print("p = ", p)
    p = p +1
until
    p == 100

print("p is ", p)

-- break 条件满足退出循环
while (p > 0) do
    goto AFTER
end

::AFTER::
     print("goto 结束")

-- goto 标签跳转
for i = 1, 10 do
    print("i is ".. i)
    if i == 5 then
        break
    end
end

-- 无限循环
while true do
    print("游戏模式已开启...")
    os.execute("sleep 10")
    print("游戏结束")
    break
end