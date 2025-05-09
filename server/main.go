package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goferwplynie/airQuality/server/internal/handlers"
	"github.com/goferwplynie/airQuality/server/internal/repository"
	"github.com/goferwplynie/airQuality/server/internal/services"
)

func main() {
	repo := repository.New()
	blayer := services.New(repo)
	handler := handlers.New(blayer)

	router := gin.New()

	router.POST("readings", handler.SaveReadingHandler)
	router.GET("readings/:timestamp", handler.GetReading)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("server running on :8080")
	}
}
