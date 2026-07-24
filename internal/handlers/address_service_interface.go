package handlers

import "example.com/address-weather-project/internal/domain"

type AddressService interface {
	FetchAddress(postalCode string) (*domain.Address, error)
}
