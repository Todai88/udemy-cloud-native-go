package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4001"
	}
	return ":" + port
}
