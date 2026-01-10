package main

import (
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus"
)

var (
    reqCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        }, []string{"method", "path", "code"},
    )
    reqLatency = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_latency_seconds",
            Help:    "HTTP request latency in seconds",
            Buckets: prometheus.DefBuckets,
        }, []string{"method", "path"},
    )
)

func init() {
    prometheus.MustRegister(reqCounter, reqLatency)
}

func prometheusMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        latency := time.Since(start).Seconds()
        reqLatency.WithLabelValues(c.Request.Method, c.FullPath()).Observe(latency)
        reqCounter.WithLabelValues(c.Request.Method, c.FullPath(), strconv.Itoa(c.Writer.Status())).Inc()
    }
}
