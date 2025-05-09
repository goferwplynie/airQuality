package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goferwplynie/airQuality/server/internal/models"
	"github.com/goferwplynie/airQuality/server/internal/services"
)

type Handler struct {
	BLayer *services.BuisnessLayer
}

func New(blayer *services.BuisnessLayer) *Handler {
	return &Handler{
		BLayer: blayer,
	}
}

func (h *Handler) SaveReadingHandler(ctx *gin.Context) {
	var request []models.PostReading

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cant read request body",
		})
	}

	errors := h.BLayer.SaveReadings(request)

	if len(errors) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		fmt.Println(errors)
	} else {
		ctx.Status(http.StatusCreated)
		fmt.Println("success")
	}
}
