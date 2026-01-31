package main

import (
	"log"
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

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Falha ao iniciar o servidor:", err)
	}
}
