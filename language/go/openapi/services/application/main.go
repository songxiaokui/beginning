package main

import (
	"fmt"
	"log"
	"opneapi/sdk"
)

func main() {
	client := sdk.Client{
		AppID:     "56b2084f-c1da-4b17-87ba-164a1e9e3686",             // 替换为服务端生成的 AppID
		SecretKey: "M2QyM2VjMWMtNmNmNy00MDdlLTg3Y2MtYzg1MzhkOGMxYTZj", // 替换为服务端生成的 SecretKey
		BaseURL:   "http://127.0.0.1:8080",
	}

	// 提交作业 (GET)
	response, err := client.SubmitJob()
	if err != nil {
		log.Fatal("SubmitJob failed:", err)
	}
	fmt.Println("Submit Response (GET):", response)

	// 取消作业 (GET)
	response, err = client.CancelJob()
	if err != nil {
		log.Fatal("CancelJob failed:", err)
	}
	fmt.Println("Cancel Response (GET):", response)

	// 提交作业 (POST)
	body := []byte(`{"job_id": "12345"}`)
	response, err = client.SendRequest("POST", "/api/submit", body)
	if err != nil {
		log.Fatal("SubmitJob POST failed:", err)
	}
	fmt.Println("Submit Response (POST):", response)
}
