package config

import (
	"log"
	"os"
)

type AppConfig struct {
	Port string
}

func NewAppConfig() *AppConfig {
	port := getEnvPort("APP_PORT")
	return &AppConfig{
		Port: port,
	}
}

func getEnvPort(key string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	log.Fatalf("Environment variable %s is not set", key)
	return ""
}
