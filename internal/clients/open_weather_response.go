package clients

import "example.com/address-weather-project/internal/domain"

type OpenWeatherResponse struct {
	Weather []OpenWeatherDescription `json:"weather"`
	Main    OpenWeatherMain          `json:"main"`
	Sys     OpenWeatherSys           `json:"sys"`
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
		Sunrise:     resp.Sys.Sunrise,
		Sunset:      resp.Sys.Sunset,
	}
}
