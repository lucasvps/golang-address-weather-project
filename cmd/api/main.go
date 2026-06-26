package main

import (
	"log"
	"net/http"
	"os"

	"example.com/address-weather-project/internal/clients"
	"example.com/address-weather-project/internal/handlers"
	"example.com/address-weather-project/internal/services"
	"example.com/address-weather-project/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := gin.Default()

	addressClient := clients.NewAddressClient(http.DefaultClient, os.Getenv("VIA_CEP_BASE_URL"))
	weatherClient := clients.NewWeatherClient()
	geocodingClient := clients.NewGeocodingClient()

	wService := services.NewWeatherService(addressClient, weatherClient, geocodingClient)

	wHandler := handlers.NewWeatherHandler(wService)

	routes.RegisterRoutes(server, wHandler)

	err = server.Run(":8080")

	if err != nil {
		panic("Ocorreu um erro ao iniciar o servidor.")
	}
}

//Get "viacep.com.br/ws/85035000/json": unsupported protocol scheme ""
