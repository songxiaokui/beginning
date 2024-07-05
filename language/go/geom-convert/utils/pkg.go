package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
)

// GenerateSign 生成签名的函数
func GenerateSign(submitUrl string, appkey string) string {
	// 将appkey添加到参数中
	params := url.Values{}
	params.Add("appkey", appkey)
	fullUrl := submitUrl + "&" + params.Encode()
	fmt.Println("finally url: ", fullUrl)
	// 计算MD5哈希值
	hash := md5.Sum([]byte(fullUrl))

	// 将哈希值转换为十六进制字符串
	sign := hex.EncodeToString(hash[:])

	return sign
}
