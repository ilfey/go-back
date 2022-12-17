package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ilfey/go-back/internal/app/config"
	"github.com/ilfey/go-back/internal/app/server"
	_ "github.com/mattn/go-sqlite3"

	"github.com/ilfey/go-back/internal/pkg/store"
	"github.com/ilfey/go-back/internal/pkg/store/pgsql"
	"github.com/ilfey/go-back/internal/pkg/store/sqlite"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	logLevel     string
	address      string
	port         string
	databaseUrl  string
	databaseFile string
	key          string
	lifeSpan     int
)

func main() {
	godotenv.Load()

	flag.StringVar(&databaseUrl, "du", getEnv("DATABASE_URL", "postgresql://ilfey:QWEasd123@localhost:5432/go-back"), "PostgreSQL database url")
	flag.StringVar(&databaseFile, "df", getEnv("DATABASE_FILE", "go-back.db"), "SQLite database url")
	flag.StringVar(&logLevel, "ll", getEnv("LOGLEVEL", "info"), "LogLevel")
	flag.StringVar(&address, "a", getEnv("ADDRESS", "0.0.0.0"), "Address")
	flag.StringVar(&port, "p", getEnv("PORT", "8000"), "Port")
	flag.StringVar(&key, "jk", getEnv("JWT_KEY", "secret"), "JWT key")
	flag.IntVar(&lifeSpan, "jls", getEnvInt("JWT_LIFE_SPAN", 24), "JWT life span (in hours)")

	flag.Parse()

	config := &config.Config{
		Address:      address + ":" + port,
		LogLevel:     logLevel,
		DatabaseUrl:  databaseUrl,
		DatabaseFile: databaseFile,
		Key:          []byte(key),
		LifeSpan:     lifeSpan,
	}

	// create logger
	logger, err := CreateLogger(config)
	if err != nil {
		logrus.Panicf("logger configuration error: %s", err.Error())
	}

	var store *store.Store
	store, err = initPgsql(config, logger)
	if err != nil {
		logger.Warn("unable to connect to postgresql database")
		store, err = initSqlite(config, logger)
		if err != nil {
			logger.Error("unable to connect to sqlite database")
		}
	}

	s := server.New()

	if err := s.Start(config, store, logger); err != nil {
		log.Fatal(err)
	}
}

func initPgsql(config *config.Config, logger *logrus.Logger) (*store.Store, error) {
	// create database connection
	db, err := pgx.Connect(context.Background(), config.DatabaseUrl)
	if err != nil {
		return nil, err
	}
	logger.Info("server connected to postgresql db")

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) UNIQUE NOT NULL,
		is_deleted BOOLEAN NOT NULL DEFAULT FALSE
	);
	`
	_, err = db.Exec(context.Background(), sqlStmt)
	if err != nil {
		return nil, err
	}
	// create store
	store := pgsql.New(db, logger)

	return store, nil
}

func initSqlite(config *config.Config, logger *logrus.Logger) (*store.Store, error) {
	// create database connection
	db, err := sql.Open("sqlite3", config.DatabaseFile)
	if err != nil {
		return nil, err
	}
	logger.Info("server connected to sqlite db")

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) UNIQUE NOT NULL,
		is_deleted BOOLEAN NOT NULL DEFAULT FALSE
	);
	`
	_, err = db.ExecContext(context.Background(), sqlStmt)
	if err != nil {
		return nil, err
	}

	// create store
	store := sqlite.New(db, logger)

	return store, nil
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
