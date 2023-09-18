package calculatecontroller

import (
	"api-akar-kuadrat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(context *gin.Context)  {
	var inputData struct {
        Number float64 `json:"number" binding:"required"`
    }

    if err := context.ShouldBindJSON(&inputData); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
	// validation
	if inputData.Number <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Value must be a positive number"})
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