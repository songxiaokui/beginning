package main

import (
	"fmt"
	"log"
	"opneapi/sdk"
)

func main() {
	client := sdk.Client{
		AppID:     "738d741b-9007-47ad-9c88-3b95e6580c65",
		SecretKey: "MDE0YjgzZjYtNDZmMC00NWQwLTk1NTYtMjU3ZTlmMmYyMDE2",
		BaseURL:   "http://127.0.0.1:8080",
	}

	// 提交作业
	response, err := client.SubmitJob()
	if err != nil {
		log.Fatal("SubmitJob failed:", err)
	}
	fmt.Println("Submit Response:", response)

	// 取消作业
	response, err = client.CancelJob()
	if err != nil {
		log.Fatal("CancelJob failed:", err)
	}
	fmt.Println("Cancel Response:", response)
}
