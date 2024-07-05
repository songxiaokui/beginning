package main

import (
	"fmt"
	"geom-convert/utils"
	"github.com/joho/godotenv"
	"net/url"
	"os"
)

var (
	APPID  string
	APPKEY string
)

func init() {
	err := godotenv.Load("./config.env")
	if err != nil {
		panic(err)
	}
	APPID = os.Getenv("APPID")
	APPKEY = os.Getenv("APPKEY")
}

func QueryResult(qid string) {
	// 基础URL
	baseUrl := "https://open.3dwhere.com/api/query"
	// URL参数
	params := url.Values{}
	params.Add("appid", APPID)
	params.Add("fileid", qid)

	submitUrl := baseUrl + "?" + params.Encode()
	// appkey
	appkey := APPKEY

	// 生成签名
	sign := utils.GenerateSign(submitUrl, appkey)

	// 将签名添加到URL参数中
	params.Add("sign", sign)

	// 构建最终的URL
	finalUrl := baseUrl + "?" + params.Encode()
	fmt.Println(finalUrl)
}

func main() {
	QueryResult("1e0Rc843RPeaXX4A")
}
