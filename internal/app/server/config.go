package server

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Address  string
	LogLevel string
}

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		Address:  getEnv("ADDRESS", ":8000"),
		LogLevel: getEnv("LOGLEVEL", "debug"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
