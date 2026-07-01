package domain

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
	Sunrise     int64   `json:"sunrise"`
	Sunset      int64   `json:"sunset"`
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
