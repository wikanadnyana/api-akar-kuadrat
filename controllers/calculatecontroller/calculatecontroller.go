package calculatecontroller

import (
	"net/http"

	"api-akar-kuadrat/models"

	"github.com/gin-gonic/gin"
)

func Create(context *gin.Context)  {
	var inputData struct {
        Number float64 `json:"number"`
    }

	// validation

    if err := context.ShouldBindJSON(&inputData); err != nil {
        context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    hasilAkarKuadrat := hitungakar(inputData.Number)

    calculation := models.Calculate{
        Number:       inputData.Number,
        Akarkuadrat: hasilAkarKuadrat,
    }

    if err := models.DB.Create(&calculation).Error; err != nil {
        context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create calculation record"})
        return
    }

    context.JSON(http.StatusOK, gin.H{"calculation": calculation.Akarkuadrat})
}

func hitungakar(number float64) float64 {
	var squareroot float64 = number / 2
	var temp float64
	for {
		temp = squareroot
		squareroot = (temp + (number / temp)) / 2
		if (temp - squareroot) == 0 {
			break
		}
	}
	return squareroot
}