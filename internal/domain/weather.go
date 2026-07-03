package domain

import "fmt"

type WeatherResponse struct {
	Weather Weather `json:"weather"`
	Address Address `json:"address"`
}

type Weather struct {
	Temperature float64 `json:"temperature"`
	FeelsLike   float64 `json:"feels_like"`
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Humidity    int     `json:"humidity"`
	Sunrise     string  `json:"sunrise"`
	Sunset      string  `json:"sunset"`
}

type Address struct {
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	State      string `json:"uf"`
	Street     string `json:"street_name"`
}

type Localization struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (wR *WeatherResponse) ToString() string {
	return fmt.Sprintf(`
		Cidade: %s
		Rua: %s
		Temperatura: %.1f°C
		Sensação térmica: %.1f°C
		Mínima: %.1f°C
		Máxima: %.1f°C
		Umidade: %d%%
		Nascer do sol: %s
		Pôr do sol: %s`,
		wR.Address.City,
		wR.Address.Street,
		wR.Weather.Temperature,
		wR.Weather.FeelsLike,
		wR.Weather.TempMin,
		wR.Weather.TempMax,
		wR.Weather.Humidity,
		wR.Weather.Sunrise,
		wR.Weather.Sunset,
	)
}
