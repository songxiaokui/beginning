-- 引入
local algorithm = require("pkg.utility")

a = 100
b = 100
print("a + b = "..algorithm.operator("+")(a, b))
print("a - b = "..algorithm.operator("-")(a, b))
print("a * b = "..algorithm.operator("*")(a, b))
print("a / b = "..algorithm.operator("/")(a, b))
print("a % b = "..algorithm.operator("%")(a, b))