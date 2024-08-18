package startFunc

import (
	"log"
	"net/http"

	"github.com/benodiwal/gorm-studio/pkg/database"
	"github.com/benodiwal/gorm-studio/pkg/env"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	env.Load()
	r := gin.Default()

	logger := log.Default()
	db := database.Connect(logger)
	println(db)
	
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H {
			"message": "pong",
		})
	})

	r.Run()
}
