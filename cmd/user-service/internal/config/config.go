package config

import (
	"os"
)

type AppConfig struct {
	Port string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Port: getEnvPort("APP_PORT", "8080"),
	}
}

func getEnvPort(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}
