--[[
Lua数据类型:
    Lua是动态语言，不需要指定变量的类型，只进行赋值即可。

8种类型:
    1. nil
    空类型 表示值无效
    2. boolean
    布尔类型 只有 true 和 false
    3. number
    浮点型 双精度浮点数
    4. string
    字符串类型 由一对单引号或双引号包围
    5. function
    函数类型 由 C 或者Lua 编写的函数
    6. userdata
    表示任意存储在变量中的 C数据结构
    7. thread
    表示执行的独立线路 和 go 等协程类似
    8. table
    表，是一个关联数组，数组的索引可以是 int 下标、string或者表类型 通过构造表达式进行 {}来创造空表

类型判断:
    type(变量名)

nil:
    删除元素
    type(不存在的变量名) 返回值为字符串"nil" 判断注意

boolean:
    Lua中只有 nil 和 false 是假
    其他为真

number:
    默认为双精度类型 double

string:
    由单引号和双引号扩起来
    由两个中括号括起来的字符串
    在进行数字字符串进行运算时，会自动进行类型转换，当然可能会转换出错
    字符串拼接使用 ..
    计算字符串长度: #变量名

table(表):
    通过{}进行表的构造

function:
    函数是一等公民
    可以将函数作为参数进行传递
    可以实现闭包
    可以进行匿名函数的使用
    定义方式: function 函数名(参数列表...)
             函数体
             end
    匿名函数: function(参数列表...)
             函数体
             end

thread:
    多线程 类似 coroutine
    和 go 的协程 goroutine 类似

ffi模块的安装与使用:
    1. 安装 brew install luarocks
    2. git clone https://github.com/facebook/luaffifb
    3. cd luaffifb; luarocks make
--]]

-- 拓展 lua 的包路径
-- package.cpath
-- nil
a = nil
if a == nil then
    print("a is nil")
else
    print("a is not nil")
end
print("a 的类型: ", type(a))

-- 注意 type(a)返回值是 nil 字符串，不要使用 nil 比较
if type(a) == "nil" then
    print("a 的类型是: ", nil)
end

-- nil 可以作为删除功能
table1 = { "五", a = "壹", b = "2", "三", "四" }
for k, v in pairs(table1) do
    print(string.format("key: %s, value: %s", k, v))
end
-- 注意: table 可以存数组和 map 键值对
-- 如果是数组元素，则从第一个开始搜索的 key 为 1、2、3，注意不是 0 开始的
-- 如果是键值对，则 k为键值

-- 删除索引为 2 的元素
-- 删除 key 为 a 的元素
print("-------------")
table1.a = nil
table1[2] = nil
for k, v in pairs(table1) do
    print(string.format("key: %s, value: %s", k, v))
end
-- 注意: 删除后 其他元素所在的索引位置是不变的

-- boolean
if false or nil then
    print("false he nil 都是假")
elseif 0 then
    print("0 为真")
elseif 1 then
    print("1 为真")
end

-- number
print(type(1))
print(type(1.1))
print(type(0.1e-23))
print(type(1e20))

-- string
s1 = "你好 世界"
s2 = '你好 时间'
print(s1)
print(s2)
s3 = [[
你好
世界
]]
print(s3)
print(1 + 1)
print("1" + 1)
print("-1" * "2")
print("2E-3" * 4)
-- print("1"+"err") 不能直接进行字符串加法
-- 字符串拼接使用 .. 进行连接
print("1" .. "err")
print(string.format("字符串 s3的长度为: %d", #s3))

-- table
-- 关联数组
local t1 = { 5, 7, 9, 11, 13 }
for k, v in pairs(t1) do
    print(k .. "->" .. v)
end
-- 遍历数组 数组所用从 1 开始，#t1 获取表的长度
for i = 1, #t1 do
    print(t1[i])
end
-- 定义映射
local t2 = { 100, 200 }
name = "sxk"
age = 10
t2[name] = age
for k, v in pairs(t2) do
    print(k .. "->" .. v)
end
print("t2表的长度: " .. #t2)

-- function
-- 函数定义
function add(sign, a, b)
    if sign == "+" then
        return a + b
    elseif sign == "-" then
        return a - b
    elseif sign == "*" then
        return a * b
    elseif sign == "/" then
        return a / b
    elseif sign == "%" then
        return a % b
    end
end

a = 20
b = 10
print("a+b=" .. add("+", a, b))
print("a-b=" .. add("-", a, b))
print("a*b=" .. add("*", a, b))
print("a/b=" .. add("/", a, b))
print("a%b=" .. add("%", a, b))

-- 函数作为参数赋值
add1 = add
print("a+b=" .. add1("+", a, b))

-- 函数作为函数参数返回
function middleware(func)
    print("执行 func 前做点什么...")
    return function()
        -- 执行函数
        func()
        -- 执行结束做点什么
        print("执行 func 后做点什么...")
        return
    end
end

function sayHello()
    print("say hello 的主要业务逻辑执行中...")
end

-- 函数作为参数进行调用
middleware(sayHello)()

-- 匿名函数
middleware(
        function()
            print("我是匿名函数的调用")
        end
)()

-- thread
-- 主要用的是 c 的 coroutine实现的
local co = coroutine.create(
        function(n)
            for i = 1, n do
                os.execute("sleep" .. " 0.1")
                print("正在执行第 " .. i .. "步")
                coroutine.yield()
            end
            print("coroutine 执行结束")
        end
)
coroutine.resume(co, 5)
coroutine.status(co)
coroutine.resume(co, 5)
coroutine.resume(co, 5)
coroutine.resume(co, 5)
coroutine.resume(co, 5)
coroutine.resume(co, 5)

-- 定义 userdata 类型使用
-- 创建一个自定义用户
-- 引入 ffi local ffi = require("ffi")
-- 需要引入: Luajit brew install luajit
local ffi = require("ffi")
-- 自定义 struct
ffi.cdef [[
struct Person
{
    int age;
    char name[3];
}
]]

local xm = ffi.new("struct Person")
xm.age = 18
xm.name = 'sxk'
-- c对象存在指针 需要释放 超过 lua 的管理范围
ffi.gc(xm, function(x)
    ffi.C.free(x.name)
    ffi.C.free(x)
end)

print(string.format("姓名:%s 年龄: %d\n", ffi.string(xm.name), xm.age))
