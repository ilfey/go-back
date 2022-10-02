package server

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Address  string
	LogLevel string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

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
