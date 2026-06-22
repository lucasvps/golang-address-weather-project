package main

import (
	"example.com/address-weather-project/internal/handlers"
	"example.com/address-weather-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	wHandler := handlers.NewWeatherHandler()

	routes.RegisterRoutes(server, wHandler)

	err := server.Run(":8080")

	if err != nil {
		panic("Ocorreu um erro ao iniciar o servidor.")
	}
}
