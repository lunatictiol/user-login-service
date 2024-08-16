package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBHost:     getEnv("DATABASE_HOST", "http://localhost"),
		DBUser:     getEnv("DATABASE_USER", "root"),
		DBPassword: getEnv("DATABASE_PASSWORD", "sabiq1234"),
		DBName:     getEnv("DATABASE_NAME", "users"),
	}
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
