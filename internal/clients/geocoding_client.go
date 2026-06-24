package clients

import "example.com/address-weather-project/internal/domain"

type GeocodingClient struct {
}

func NewGeocodingClient() *GeocodingClient {
	return &GeocodingClient{}
}

func (c *GeocodingClient) FetchGeocoding(address domain.Address) (domain.Localization, error) {
	return domain.Localization{
		Latitude:  "valor-da-latitude-aqui",
		Longitude: "valor da longitude aqui",
	}, nil
}
