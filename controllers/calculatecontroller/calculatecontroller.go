package calculatecontroller

import (
	"math"
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

    hasilAkarKuadrat := math.Sqrt(inputData.Number)

    calculation := models.Calculate{
        Number:       inputData.Number,
        HasilAkarKuadrat: hasilAkarKuadrat,
    }

    if err := models.DB.Create(&calculation).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create calculation record"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"calculation": calculation})
}