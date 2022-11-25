package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ilfey/go-back/internal/app/config"
	"github.com/ilfey/go-back/internal/app/server"
	"github.com/joho/godotenv"
)

var (
	logLevel    string
	address     string
	port        string
	databaseUrl string
	key         string
	lifeSpan    int
)

func main() {
	godotenv.Load()

	flag.StringVar(&databaseUrl, "du", getEnv("DATABASE_URL", "PostgreSQL database url"), "LogLevel")
	flag.StringVar(&logLevel, "ll", getEnv("LOGLEVEL", "info"), "LogLevel")
	flag.StringVar(&address, "a", getEnv("ADDRESS", "0.0.0.0"), "Address")
	flag.StringVar(&port, "p", getEnv("PORT", "8000"), "Port")
	flag.StringVar(&key, "jk", getEnv("JWT_KEY", "secret"), "JWT key")
	flag.IntVar(&lifeSpan, "jls", getEnvInt("JWT_LIFE_SPAN", 24), "JWT life span (in hours)")

	flag.Parse()

	config := &config.Config{
		Address:     address + ":" + port,
		LogLevel:    logLevel,
		DatabaseUrl: databaseUrl,
		Key:         []byte(key),
		LifeSpan:    lifeSpan,
	}

	s := server.New()

	if err := s.Start(config); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if s, ok := os.LookupEnv(key); ok {
		value, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		return value
	}
	return fallback
}
