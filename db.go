package main

import (
    "context"
    "log"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
)

func connectDB(cfg *Config) (*pgxpool.Pool, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    pool, err := pgxpool.New(ctx, cfg.DBDSN)
    if err != nil {
        return nil, err
    }
    // try ping
    if err := pool.Ping(ctx); err != nil {
        log.Printf("db ping failed: %v", err)
        // return pool anyway; readiness will fail until db responsive
    }
    return pool, nil
}
