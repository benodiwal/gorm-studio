package routes

import "github.com/gin-contrib/cors"

func (r *Router) RegisterMiddlewares() {
	r.Engine.Use(cors.Default())
}
