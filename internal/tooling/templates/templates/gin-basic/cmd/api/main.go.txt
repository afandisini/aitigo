package main

import (
	"log"
	"net/http"
	"os"

	transportgin "example.com/aitigo-gin-basic/internal/transport/http/gin"
	"example.com/aitigo-gin-basic/internal/transport/http/gin/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())
	transportgin.Register(router)

	addr := ":" + port
	if err := router.Run(addr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
