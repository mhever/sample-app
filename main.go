package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg := loadConfig()

	db, err := connectDB(cfg)
	if err != nil {
		log.Printf("warning: db connect failed: %v", err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(prometheusMiddleware())

	// basic endpoints
	r.GET("/", func(c *gin.Context) { handleRoot(c, cfg) })
	r.GET("/healthz", handleHealth)
	r.GET("/readyz", func(c *gin.Context) { handleReady(c, db) })
	r.GET("/env", func(c *gin.Context) { handleEnv(c, cfg) })
	r.GET("/work", handleWork)
	r.GET("/kill", handleKill)

	// metrics endpoint (Prometheus)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	addr := ":8080"
	log.Printf("starting %s on %s", cfg.ServiceName, addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}

	// ensure DB closed on exit
	if db != nil {
		db.Close()
	}
	// keep the compiler happy about context import in simple app
	_ = context.TODO()
}
