package config

import (
	"fmt"
	"log"
	"os"
	"user-service/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SSLMode  string
}

func InitDB() {
	cfg := DBConfig{
		Host:     getEnv("DB_HOST"),
		User:     getEnv("DB_USER"),
		Password: getEnv("DB_PASSWORD"),
		Name:     getEnv("DB_NAME"),
		Port:     getEnv("DB_PORT"),
		SSLMode:  getEnv("DB_SSLMODE"),
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	DB = db
	log.Println("Успешное подключение к базе")
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Необходимо установить переменную окружения: %s", key)
	}
	return value
}
