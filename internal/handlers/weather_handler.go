package handlers

import (
	"net/http"

	"example.com/address-weather-project/internal/services"
	"example.com/address-weather-project/internal/validation"
	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	wService *services.WeatherService
}

func NewWeatherHandler(wService *services.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		wService: wService,
	}
}

func (h *WeatherHandler) FetchWeatherByPostalCode(context *gin.Context) {
	postalCode := context.Param("postalCode")

	if !validation.IsPostalCodeValid(postalCode) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "The postalCode is invalid"})
		return
	}

	data, err := h.wService.FetchWeatherDataFromPostalCode(postalCode)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error ocurred."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": data})
}
