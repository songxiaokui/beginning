package main

import (
	"fmt"
	"geom-convert/utils"
	"github.com/joho/godotenv"
	"net/url"
	"os"
)

func init() {
	err := godotenv.Load("./config.env")
	if err != nil {
		panic(err)
	}
	APPID = os.Getenv("APPID")
	APPKEY = os.Getenv("APPKEY")
}

func main1() {
	// 基础URL
	baseUrl := "https://open.3dwhere.com/api/add"
	// URL参数
	params := url.Values{}
	params.Add("appid", APPID)
	params.Add("infile", "https://img.austsxk.com/austsxk/huojia.x_t")
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
