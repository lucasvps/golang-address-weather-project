package routes

import (
	"example.com/address-weather-project/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, wHandler *handlers.WeatherHandler, aHandler *handlers.AddressHandler) {
	server.GET("/weather/:postalCode", wHandler.FetchWeatherByPostalCode)
	server.GET("/address/:postalCode", aHandler.FetchAddressFromPostalCode)
}
