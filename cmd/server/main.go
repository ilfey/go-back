package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ilfey/go-back/internal/app/config"
	"github.com/ilfey/go-back/internal/app/server"
	"github.com/ilfey/go-back/internal/pkg/store/pgsql"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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

	// create logger
	logger, err := CreateLogger(config)
	if err != nil {
		logrus.Panicf("logger configuration error: %s", err.Error())
	}

	// create database connection
	db, err := pgx.Connect(context.Background(), config.DatabaseUrl)
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("server connected to db")
	}

	// create store
	store := pgsql.New(db, logger)

	s := server.New()

	if err := s.Start(config, store, logger); err != nil {
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

func CreateLogger(config *config.Config) (*logrus.Logger, error) {
	logger := logrus.New()
	level, err := logrus.ParseLevel(config.LogLevel)

	if err != nil {
		return nil, err
	}

	logger.SetLevel(level)

	return logger, nil
}
