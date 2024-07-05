package main

import (
	"fmt"
	"geom-convert/utils"
	"net/url"
)

func main() {
	// 基础URL
	baseUrl := "https://open.3dwhere.com/api/add"
	// URL参数
	params := url.Values{}
	params.Add("appid", APPID)
	params.Add("infile", "https://img.austsxk.com/austsxk/cx.x_t")
	params.Add("outtype", "stp")

	submitUrl := baseUrl + "?" + params.Encode()
	// appkey
	appkey := APPKEY

	// 生成签名
	sign := utils.GenerateSign(submitUrl, appkey)

	// 将签名添加到URL参数中
	params.Add("sign", sign)

	// 构建最终的URL
	finalUrl := baseUrl + "?" + params.Encode()

	// 输出最终的请求URL
	fmt.Println(finalUrl)
}
