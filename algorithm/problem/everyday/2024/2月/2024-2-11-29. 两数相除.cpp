/*
29. 两数相除
中等
相关标签
相关企业
给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。

整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。

返回被除数 dividend 除以除数 divisor 得到的 商 。

注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−231,  231 − 1] 。本题中，如果商 严格大于 231 − 1 ，则返回 231 − 1 ；如果商 严格小于 -231 ，则返回 -231 。



示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = 3.33333.. ，向零截断后得到 3 。
*/

/*
暴力法:
 被除数 - 除数, 循环操作, 直到被除数< 0, 统计操作的次数即可，但是会超时
*/

func divide(dividend int, divisor int) int {
    var count = 0
    var sign = 1
    if dividend > 0 && divisor > 0 {
        sign = 1
    } else if (dividend * divisor) > 0 {
        if dividend < 0 && divisor < 0 {
            sign = 1
            dividend = -dividend
            divisor = -divisor
        }
    } else {
        sign = -1
        if dividend > 0 {
            divisor = -divisor
        } else {
            dividend = -dividend
        }
    }

    for dividend - divisor >= 0 {
        count++;
        dividend -=  divisor
    }

    result := sign*count

    // 处理边界问题
    if result > math.MaxInt32 {
        return math.MaxInt32
    }
    if result < math.MinInt32 {
        return math.MinInt32
    }
    return  result
}

// 位运算
func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }

    sign := 1
    if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
        sign = -1
    }

    dividend = abs(dividend)
    divisor = abs(divisor)

    result := 0
    for dividend >= divisor {
        temp, multiple := divisor, 1
        for dividend >= (temp << 1) {
            temp <<= 1
            multiple <<= 1
        }
        dividend -= temp
        result += multiple
    }

    return sign * result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}