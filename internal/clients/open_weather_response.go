package clients

import (
	"time"

	"example.com/address-weather-project/internal/domain"
)

type OpenWeatherResponse struct {
	Weather  []OpenWeatherDescription `json:"weather"`
	Main     OpenWeatherMain          `json:"main"`
	Sys      OpenWeatherSys           `json:"sys"`
	Timezone int                      `json:"timezone"`
}

type OpenWeatherDescription struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type OpenWeatherMain struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Humidity  int     `json:"humidity"`
}

type OpenWeatherSys struct {
	Sunrise int64 `json:"sunrise"`
	Sunset  int64 `json:"sunset"`
}

func (resp *OpenWeatherResponse) ToWeatherDomain() *domain.Weather {
	return &domain.Weather{
		Temperature: resp.Main.Temp,
		FeelsLike:   resp.Main.FeelsLike,
		TempMin:     resp.Main.TempMin,
		TempMax:     resp.Main.TempMax,
		Description: resp.Weather[0].Description,
		Category:    resp.Weather[0].Main,
		Humidity:    resp.Main.Humidity,
		Sunrise:     formatUnixTime(resp.Sys.Sunrise, resp.Timezone),
		Sunset:      formatUnixTime(resp.Sys.Sunset, resp.Timezone),
	}
}

func formatUnixTime(timestamp int64, timezoneOffset int) string {
	location := time.FixedZone("openweather-location", timezoneOffset)
	return time.Unix(timestamp, 0).In(location).Format("15:04")
}
