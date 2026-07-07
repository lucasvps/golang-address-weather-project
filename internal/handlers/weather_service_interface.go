package handlers

import "example.com/address-weather-project/internal/domain"

type WeatherService interface {
	FetchWeatherDataFromPostalCode(postalCode string) (*domain.WeatherResponse, error)
}
