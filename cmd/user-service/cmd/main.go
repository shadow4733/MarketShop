package main

import (
	"log"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/router"
	"user-service/internal/service/impl"

	_ "user-service/cmd/docs"
)

// @title User Service API
// @version 1.0
// @description REST API для сервиса пользователей

// @contact.name API Support
// @contact.email support@userservice.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

func main() {
	config.InitDB()
	appCfg := config.NewAppConfig()

	authService := impl.NewAuthService(config.DB)
	authHandler := handler.NewAuthHandler(authService)
	ginRouter := router.SetupRouter(authHandler)

	port := appCfg.Port

	log.Printf("Документация Swagger: http://localhost:%s/swagger/index.html", port)

	if err := ginRouter.Run(":" + port); err != nil {
		log.Fatal("Ошибка подключения к серверу: ", err)
	}
}
