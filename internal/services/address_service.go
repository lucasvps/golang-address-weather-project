package services

import (
	"log/slog"

	"example.com/address-weather-project/internal/domain"
)

type AddressService struct {
	clientAddress AddressClient
	logger        *slog.Logger
}

func NewAddressService(clientAddress AddressClient, logger *slog.Logger) *AddressService {
	return &AddressService{
		clientAddress: clientAddress,
		logger:        logger,
	}
}

func (aService AddressService) FetchAddress(postalCode string) (*domain.Address, error) {
	aService.logger.Info("started fetch address flow", "postal_code", postalCode)

	address, err := aService.clientAddress.FetchAddress(postalCode)

	if err != nil {
		aService.logger.Error("failed to fetch address", "postal_code", postalCode, "error", err)
		return nil, err
	}

	aService.logger.Info("address fetched successfully", "postal_code", postalCode)

	return &address, nil
}
