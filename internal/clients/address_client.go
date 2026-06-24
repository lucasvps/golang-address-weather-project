package clients

import "example.com/address-weather-project/internal/domain"

type AddressClient struct {
}

func NewAddressClient() *AddressClient {
	return &AddressClient{}
}

func (c *AddressClient) FetchAddress(postalCode string) (domain.Address, error) {
	return domain.Address{
		PostalCode: postalCode,
		City:       "Guarapuava",
		State:      "Parana",
		Street:     "Rua São Paulo",
	}, nil
}
