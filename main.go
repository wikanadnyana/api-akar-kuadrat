package main

import (
	"api-akar-kuadrat/controllers/calculatecontroller"
	"api-akar-kuadrat/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	router.POST("/api/calculate", calculatecontroller.Create)

	router.Run()
}
