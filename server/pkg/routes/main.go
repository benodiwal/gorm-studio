package routes

import (
	"github.com/benodiwal/gorm-studio/pkg/database"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	database *database.Database
}

func New(database *database.Database) *Router {
	return &Router{
		Engine: gin.New(),
		database: database,
	}
}
