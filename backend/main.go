package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"myapp/backend/routes"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	routes.Register(r)
	r.Run(":8080")
}
