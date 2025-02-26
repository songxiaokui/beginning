package main

import (
	"deepseek_access/internal/entity"
	"deepseek_access/internal/handler"
	"deepseek_access/internal/middleware"
	"deepseek_access/internal/routes"
	"deepseek_access/internal/svc"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configPath := flag.String("config", "", "Path to the YAML configuration file")
	flag.Parse()

	_config, err := entity.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	services := svc.NewServiceContext(_config)

	chatHandler := handler.NewChatHandler(services)

	router := gin.Default()
	_ = router.SetTrustedProxies(nil)

	router.Use(middleware.Cors())

	routes.RegisterRoutes(router, chatHandler)

	log.Printf("服务器启动在 :%d \n", _config.Port)
	if err := router.Run(fmt.Sprintf(":%d", _config.Port)); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
