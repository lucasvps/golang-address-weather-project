package routes

import (
	"example.com/address-weather-project/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, wHandler *handlers.WeatherHandler) {
	server.GET("/weather/:postalCode", wHandler.FetchWeatherByPostalCode)
}
