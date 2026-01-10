package main

import (
	"fmt"
	"os"
)

type Config struct {
	ServiceName string
	Version     string
	DBDSN       string
}

func loadConfig() *Config {
	dbHost := getenv("DB_HOST", "localhost")
	dbPort := getenv("DB_PORT", "5432")
	dbUser := getenv("DB_USER", "postgres")
	dbPassword := getenv("DB_PASSWORD", "password")
	dbName := getenv("DB_NAME", "postgres")
	dbSSLMode := getenv("DB_SSLMODE", "disable")

	dbDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	return &Config{
		ServiceName: getenv("SERVICE_NAME", "sample-app"),
		Version:     getenv("VERSION", "0.1.0"),
		DBDSN:       getenv("DB_DSN", dbDSN), // Allow override with full DSN if needed
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
