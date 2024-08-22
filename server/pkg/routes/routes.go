package routes

import (
	"net/http"

	"github.com/benodiwal/gorm-studio/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func (r *Router) RegisterRoutes() {
	r.Engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Ok")
	})

	r.Engine.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Ok")
	})

	handler := handlers.New(r.database, r.Engine)
	handler.RegisterRoutes()
}
