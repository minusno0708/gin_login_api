package main

import (
	"jwt-gin/controllers"
    "jwt-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)

	router.Run(":8080")
}