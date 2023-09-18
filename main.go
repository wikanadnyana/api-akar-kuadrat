package main

import (
	"api-akar-kuadrat/controllers/calculatecontroller"
	"api-akar-kuadrat/models"

	"github.com/gin-gonic/gin"
)

func main()  {
	router:= gin.Default()
	models.ConnectDatabase()

	router.POST("/api/calculate", calculatecontroller.Create)

	router.Run()
}