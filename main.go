package main

import (
	"ecommerce-api/models"
	"ecommerce-api/routes"
	"log"
)

func main() {
	// 初始化数据库
	models.InitDB()

	// 初始化路由
	r := routes.InitRoutes()

	// 启动服务器
	log.Println("Server is running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
