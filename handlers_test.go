package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleHealth(t *testing.T) {
	// Create a Gin router
	r := gin.New()
	r.GET("/healthz", handleHealth)

	// Create a test HTTP request
	req, _ := http.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status": "ok"}`, w.Body.String())
}
