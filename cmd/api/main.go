package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"example.com/address-weather-project/internal/cache"
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

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	addressClient := clients.NewAddressClient(http.DefaultClient, os.Getenv("VIA_CEP_BASE_URL"), logger)
	weatherClient := clients.NewWeatherClient(http.DefaultClient, os.Getenv("OPEN_WEATHER_BASE_URL"), os.Getenv("OPEN_WEATHER_API_KEY"), logger)
	geocodingClient := clients.NewGeocodingClient(http.DefaultClient, os.Getenv("NOMINATIM_BASE_URL"), logger)
	weatherCache := cache.NewWeatherCache(time.Minute * 10)

	wService := services.NewWeatherService(addressClient, weatherClient, geocodingClient, weatherCache, logger)
	addressService := services.NewAddressService(addressClient, logger)

	wHandler := handlers.NewWeatherHandler(wService)

	aHandler := handlers.NewAddressHandler(addressService)

	routes.RegisterRoutes(server, wHandler, aHandler)

	err = server.Run(":8080")

	if err != nil {
		panic("Ocorreu um erro ao iniciar o servidor.")
	}
}
