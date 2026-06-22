package handlers

import (
	"net/http"

	"example.com/address-weather-project/internal/domain"
	"example.com/address-weather-project/internal/validation"
	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}

func (handler *WeatherHandler) FetchWeatherByPostalCode(context *gin.Context) {
	postalCode := context.Param("postalCode")

	if !validation.IsPostalCodeValid(postalCode) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "The postalCode is invalid"})
		return
	}

	//	TODO: Remover mock aqui e conectar com o service. O service vai compor as chamadas necessárias.

	data := domain.WeatherResponse{
		Weather: domain.Weather{
			Temperature: 6,
			Description: "Descrição aqui",
			Humidity:    80,
		},
		Address: domain.Address{
			PostalCode: postalCode,
			City:       "Guarapuava",
			State:      "Parana",
			Street:     "Rua São Paulo",
		},
	}

	context.JSON(http.StatusOK, gin.H{"data": data})
}
