package services

import (
	"fmt"

	"example.com/address-weather-project/internal/domain"
)

type WeatherService struct {
	addressClient   AddressClient
	weatherClient   WeatherClient
	geocodingClient GeocodingClient
}

func NewWeatherService(addressClient AddressClient,
	weatherClient WeatherClient,
	geocodingClient GeocodingClient) *WeatherService {
	return &WeatherService{
		addressClient:   addressClient,
		weatherClient:   weatherClient,
		geocodingClient: geocodingClient,
	}
}

func (wService *WeatherService) FetchWeatherDataFromPostalCode(postalCode string) (*domain.WeatherResponse, error) {

	addressData, err := wService.addressClient.FetchAddress(postalCode)

	if err != nil {
		return nil, err
	}

	latLongData, err := wService.geocodingClient.FetchGeocoding(addressData)

	if err != nil {
		return nil, err
	}

	weatherData, err := wService.weatherClient.FetchWeather(latLongData.Latitude, latLongData.Longitude)

	if err != nil {
		return nil, err
	}

	responseData := domain.WeatherResponse{
		Weather: weatherData,
		Address: addressData,
	}

	stringResponseData := responseData.ToString()

	fmt.Println(stringResponseData)

	return &responseData, nil
}
