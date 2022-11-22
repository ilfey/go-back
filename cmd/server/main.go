package main

import (
	"flag"
	"log"
	"os"

	"github.com/ilfey/go-back/internal/app/server"
	"github.com/joho/godotenv"
)

var (
	logLevel    string
	address     string
	port        string
	databaseUrl string
	key         string
)

func main() {
	godotenv.Load()

	flag.StringVar(&databaseUrl, "du", getEnv("DATABASE_URL", "PostgreSQL database url"), "LogLevel")
	flag.StringVar(&logLevel, "ll", getEnv("LOGLEVEL", "info"), "LogLevel")
	flag.StringVar(&address, "a", getEnv("ADDRESS", "0.0.0.0"), "Address")
	flag.StringVar(&port, "p", getEnv("PORT", "8000"), "Port")
	flag.StringVar(&key, "k", getEnv("JWT_KEY", "secret"), "JWT key")

	flag.Parse()

	config := &server.Config{
		Address:     address + ":" + port,
		LogLevel:    logLevel,
		DatabaseUrl: databaseUrl,
		Key:         []byte(key),
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
