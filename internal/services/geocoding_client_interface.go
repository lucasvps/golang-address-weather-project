package services

import "example.com/address-weather-project/internal/domain"

type GeocodingClient interface {
	FetchGeocoding(address domain.Address) (domain.Localization, error)
}
