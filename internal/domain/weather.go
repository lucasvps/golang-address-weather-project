package domain

type WeatherResponse struct {
	Weather Weather `json:"weather"`
	Address Address `json:"address"`
}

type Weather struct {
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
	Humidity    int64   `json:"humidity"`
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
