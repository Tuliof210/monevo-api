package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"system":  "Monevo API v2.0",
		})
	})

	router.Run(":8080") // Roda na porta 8080
}
