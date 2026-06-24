package services

import "example.com/address-weather-project/internal/domain"

type AddressClient interface {
	FetchAddress(postalCode string) (domain.Address, error)
}
