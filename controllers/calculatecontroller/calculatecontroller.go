package calculatecontroller

import (
	"api-akar-kuadrat/models"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(context *gin.Context) {
	var inputData struct {
		Number float64 `json:"number"`
	}

	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// validation
	if inputData.Number < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Value must be a positive number"})
		return
	}

	var hasilAkarKuadrat float64

	if inputData.Number == 0 {
		hasilAkarKuadrat = 0
	}else{
		hasilAkarKuadrat = hitungakar(inputData.Number)
	}
	


	if hasilAkarKuadrat == -92233720368547.77{
		context.JSON(http.StatusBadRequest, gin.H{"error": "Angka melebihi kapasitas"})
		return
	}


	calculation := models.Calculate{
		Number:      inputData.Number,
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
	// mengambil 5 angka belakang koma
	squareroot *= 100000
	round_squareroot := int(math.Floor(squareroot))
	squareroot = float64(round_squareroot) / 100000
	return squareroot
}
