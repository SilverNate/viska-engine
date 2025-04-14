package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN     string
	Port      string
	JWTSecret string
	RedisAddr string
	RateLimit int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DBDSN:     getEnv("DB_DSN", ""),
		Port:      getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),
		RateLimit: getEnvAsInt("RATE_LIMIT", 5),
	}
}

func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	if valueStr := os.Getenv(name); valueStr != "" {
		var val int
		_, err := fmt.Sscanf(valueStr, "%d", &val)
		if err == nil {
			return val
		}
	}
	return defaultVal
}
