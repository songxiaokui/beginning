--[[
算数包
--]]
Algorithm = {
}

function add(a, b)
    return a + b
end

function sub(a, b)
    return a - b
end

function multi(a, b)
    return a * b
end

function div(a, b)
    return a / b
end

function mod(a, b)
    return a % b
end

function Algorithm.operator(sign)
    if sign == "+" then
        return add
    elseif sign == "-" then
        return sub
    elseif sign == "*" then
        return multi
    elseif sign == "/" then
        return div
    elseif sign == "%" then
        return mod
    end
end

return Algorithm