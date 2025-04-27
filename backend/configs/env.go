package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnectionURL string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	defaultConnectionURL := "postgresql://admin:something@localhost:5432/cubeit-local?sslmode=disable"

	return Config{
		DBConnectionURL: getEnv("DB_CONNECTION_STRING", defaultConnectionURL),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
