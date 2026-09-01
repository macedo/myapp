package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"myapp/backend/handlers"
)

func Register(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Capitalos-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		api.GET("/health", handlers.Health)
		api.GET("/get-capitalos-token", handlers.GetCapitalosToken)
	}
}
