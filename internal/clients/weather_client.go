package clients

import "example.com/address-weather-project/internal/domain"

type WeatherClient struct {
}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{}
}

func (c *WeatherClient) FetchWeather(lat string, long string) (domain.Weather, error) {
	return domain.Weather{
		Temperature: 6,
		Description: "Descrição aqui",
		Humidity:    80,
	}, nil
}
