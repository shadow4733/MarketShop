package main

import (
	"log"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/router"
	"user-service/internal/service"

	_ "user-service/cmd/docs"
)

// @title User Service API
// @version 1.0
// @description REST API для сервиса пользователей

// @contact.name API Support
// @contact.email support@userservice.com

// @host localhost:8080
// @BasePath /api/v1

func main() {
	config.InitDB()
	appCfg := config.NewAppConfig()

	userService := service.NewUserServiceImpl(config.DB)
	userHandler := handler.NewUserHandler(userService)
	ginRouter := router.SetupRouter(userHandler)

	port := appCfg.Port

	log.Printf("Документация Swagger: http://localhost:%s/swagger/index.html", port)

	if err := ginRouter.Run(":" + port); err != nil {
		log.Fatal("Ошибка подключения к серверу: ", err)
	}
}
