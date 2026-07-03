package services

import (
	"log/slog"

	"example.com/address-weather-project/internal/domain"
)

type WeatherService struct {
	addressClient   AddressClient
	weatherClient   WeatherClient
	geocodingClient GeocodingClient
	logger          *slog.Logger
}

func NewWeatherService(addressClient AddressClient, weatherClient WeatherClient, geocodingClient GeocodingClient, logger *slog.Logger) *WeatherService {
	return &WeatherService{
		addressClient:   addressClient,
		weatherClient:   weatherClient,
		geocodingClient: geocodingClient,
		logger:          logger,
	}
}

func (wService *WeatherService) FetchWeatherDataFromPostalCode(postalCode string) (*domain.WeatherResponse, error) {

	wService.logger.Info("started weather flow", "postal_code", postalCode)

	addressData, err := wService.addressClient.FetchAddress(postalCode)

	if err != nil {
		wService.logger.Error("failed to fetch address", "postal_code", postalCode, "error", err)
		return nil, err
	}

	latLongData, err := wService.geocodingClient.FetchGeocoding(addressData)

	if err != nil {
		wService.logger.Error("failed to fetch geocoding", "postal_code", postalCode, "error", err)
		return nil, err
	}

	weatherData, err := wService.weatherClient.FetchWeather(latLongData.Latitude, latLongData.Longitude)

	if err != nil {
		wService.logger.Error("failed to fetch weather", "postal_code", postalCode, "error", err)
		return nil, err
	}

	responseData := domain.WeatherResponse{
		Weather: weatherData,
		Address: addressData,
	}

	wService.logger.Info("finished weather flow", "postal_code", postalCode)

	return &responseData, nil
}
