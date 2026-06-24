package services

import "example.com/address-weather-project/internal/domain"

type WeatherClient interface {
	FetchWeather(lat string, long string) (domain.Weather, error)
}
