package handlers

import (
	"net/http"

	"example.com/address-weather-project/internal/validation"
	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	addressService AddressService
}

func NewAddressHandler(addressService AddressService) *AddressHandler {
	return &AddressHandler{
		addressService: addressService,
	}
}

func (aHandler AddressHandler) FetchAddressFromPostalCode(context *gin.Context) {
	postalCodeParam := context.Param("postalCode")

	if !validation.IsPostalCodeValid(postalCodeParam) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "The postal code is invalid.", "postal_code": postalCodeParam})
		return
	}

	address, err := aHandler.addressService.FetchAddress(postalCodeParam)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": address})
}
