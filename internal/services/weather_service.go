package services

import (
	"log/slog"

	"example.com/address-weather-project/internal/cache"
	"example.com/address-weather-project/internal/domain"
)

type WeatherService struct {
	addressClient   AddressClient
	weatherClient   WeatherClient
	geocodingClient GeocodingClient
	weatherCache    *cache.WeatherCache
	logger          *slog.Logger
}

func NewWeatherService(addressClient AddressClient, weatherClient WeatherClient, geocodingClient GeocodingClient, weatherCache *cache.WeatherCache, logger *slog.Logger) *WeatherService {
	return &WeatherService{
		addressClient:   addressClient,
		weatherClient:   weatherClient,
		geocodingClient: geocodingClient,
		weatherCache:    weatherCache,
		logger:          logger,
	}
}

func (wService *WeatherService) FetchWeatherDataFromPostalCode(postalCode string) (*domain.WeatherResponse, error) {
	wService.logger.Info("started weather flow", "postal_code", postalCode)

	data := wService.weatherCache.Get(postalCode)

	if data != nil {
		wService.logger.Info("found weather data on cache", "postal_code", postalCode)
		return data, nil
	}

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

	wService.logger.Info("saving weather data on cache", "postal_code", postalCode)

	wService.weatherCache.Set(postalCode, &responseData)

	wService.logger.Info("finished weather flow", "postal_code", postalCode)

	return &responseData, nil
}
