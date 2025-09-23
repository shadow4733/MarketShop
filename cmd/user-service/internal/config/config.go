package config

import (
	"log"
	"os"
)

type AppConfig struct {
	Port string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Port: getEnvPort("APP_PORT"),
	}
}

func getEnvPort(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Необходимо установить переменную окружения: %s", key)
	}
	return value
}
