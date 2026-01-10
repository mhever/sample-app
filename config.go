package main

import (
    "os"
)

type Config struct {
    ServiceName string
    Version     string
    DBDSN       string
}

func loadConfig() *Config {
    return &Config{
        ServiceName: getenv("SERVICE_NAME", "sample-app"),
        Version:     getenv("VERSION", "0.1.0"),
        DBDSN:       getenv("DB_DSN", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"),
    }
}

func getenv(k, def string) string {
    if v := os.Getenv(k); v != "" {
        return v
    }
    return def
}
