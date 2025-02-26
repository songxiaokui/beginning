package main

import (
	"deepseek_access/internal/entity"
	"deepseek_access/internal/handler"
	"deepseek_access/internal/middleware"
	"deepseek_access/internal/routes"
	"deepseek_access/internal/svc"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	//yamlFile := `/Users/austsxk/Desktop/开端/笔记/beginning/language/go/deepseek_access/internal/etc/deepseek.yaml` // Replace with your YAML file path
	yamlFile := `/Users/austsxk/Desktop/开端/笔记/beginning/language/go/deepseek_access/internal/etc/default.yaml` // Replace with your YAML file path
	envFile := `/Users/austsxk/Desktop/开端/笔记/beginning/language/go/deepseek_access/internal/etc/.env`
	var config entity.APIConfig
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}
	config.APIKey = os.Getenv("DEEPSEEK_SK")

	services := svc.NewServiceContext(config)

	// 全量输出
	/*
		resp, err := services.LLMClient.Generate(context.Background(), "给我讲一个三国时期的笑话，关于曹操的，字数控制在300字内")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.Content)

	*/

	/*
		// 流式输出
		// 获取流式通道
		streamCh := services.LLMClient.StreamGenerate(context.Background(),
			"给我讲一个三国时期的笑话，关于曹操的，字数控制在1000字内")

		// 实时打印流式输出
		for {
			select {
			case resp, ok := <-streamCh:
				if !ok {
					return // 通道关闭
				}
				if resp.Error != nil {
					log.Printf("发生错误: %v", resp.Error)
					return
				}
				fmt.Print(resp.Content)
			}
		}

	*/
	// 初始化服务
	chatHandler := handler.NewChatHandler(services)

	// 创建Gin实例
	router := gin.Default()
	router.Use(middleware.Cors())

	// 注册路由
	routes.RegisterRoutes(router, chatHandler)

	// 启动服务器
	log.Println("服务器启动在 :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
