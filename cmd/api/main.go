package main

import (
	"example.com/address-weather-project/internal/clients"
	"example.com/address-weather-project/internal/handlers"
	"example.com/address-weather-project/internal/services"
	"example.com/address-weather-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	addressClient := clients.NewAddressClient()
	weatherClient := clients.NewWeatherClient()
	geocodingClient := clients.NewGeocodingClient()

	wService := services.NewWeatherService(addressClient, weatherClient, geocodingClient)

	wHandler := handlers.NewWeatherHandler(wService)

	routes.RegisterRoutes(server, wHandler)

	err := server.Run(":8080")

	if err != nil {
		panic("Ocorreu um erro ao iniciar o servidor.")
	}
}
