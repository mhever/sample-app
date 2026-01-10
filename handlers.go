package main

import (
    "context"
    "net/url"
    "os"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5/pgxpool"
)

func handleRoot(c *gin.Context, cfg *Config) {
    hn, _ := os.Hostname()
    c.JSON(200, gin.H{
        "service": cfg.ServiceName,
        "version": cfg.Version,
        "hostname": hn,
        "time": time.Now().UTC().Format(time.RFC3339),
    })
}

func handleHealth(c *gin.Context) {
    c.JSON(200, gin.H{"status": "ok"})
}

func handleReady(c *gin.Context, db *pgxpool.Pool) {
    if db == nil {
        c.JSON(503, gin.H{"ready": false, "reason": "no-db-configured"})
        return
    }
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    if err := db.Ping(ctx); err != nil {
        c.JSON(503, gin.H{"ready": false, "reason": err.Error()})
        return
    }
    c.JSON(200, gin.H{"ready": true})
}

func handleEnv(c *gin.Context, cfg *Config) {
    // return a small subset (safe to display)
    dbInfo := map[string]string{}
    if u, err := url.Parse(cfg.DBDSN); err == nil {
        dbInfo["scheme"] = u.Scheme
        dbInfo["host"] = u.Host
        dbInfo["path"] = u.Path
    }
    c.JSON(200, gin.H{
        "service": cfg.ServiceName,
        "version": cfg.Version,
        "db": dbInfo,
    })
}

func handleWork(c *gin.Context) {
    msStr := c.DefaultQuery("ms", "200")
    ms, _ := strconv.Atoi(msStr)
    if ms < 0 {
        ms = 0
    }
    // sleep to simulate work
    time.Sleep(time.Duration(ms) * time.Millisecond)
    c.JSON(200, gin.H{"slept_ms": ms})
}
