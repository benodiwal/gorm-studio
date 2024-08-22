package handlers

import (
	"github.com/benodiwal/gorm-studio/pkg/database"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router *gin.Engine
	database *database.Database
}

func (h *Handler) RegisterRoutes() {
	h.router.Group("/")
	h.router.GET("/models", h.GetModels)
}

func New(database *database.Database, router *gin.Engine) *Handler {
	return &Handler{
		database: database,
		router: router,
	}
}
