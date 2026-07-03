package clients

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"example.com/address-weather-project/internal/domain"
)

type AddressClient struct {
	httpClient *http.Client
	baseUrl    string
	logger     *slog.Logger
}

func NewAddressClient(httpClient *http.Client, baseUrl string, logger *slog.Logger) *AddressClient {
	return &AddressClient{
		httpClient: httpClient,
		baseUrl:    baseUrl,
		logger:     logger,
	}
}

func (c *AddressClient) FetchAddress(postalCode string) (domain.Address, error) {
	requestUrl := c.baseUrl + postalCode + "/json"

	c.logger.Info("fetching address", "postal_code", postalCode)

	resp, err := c.httpClient.Get(requestUrl)

	if err != nil {
		return domain.Address{}, err
	}

	defer resp.Body.Close()

	c.logger.Info("fetch address status", "postal_code", postalCode, "status_code", resp.StatusCode, "provider", "viacep")

	if resp.StatusCode != http.StatusOK {
		return domain.Address{}, fmt.Errorf("viacep returned non-ok status %d", resp.StatusCode)
	}

	var responseData ViaCepResponse

	err = json.NewDecoder(resp.Body).Decode(&responseData)

	if err != nil {
		c.logger.Error("error decoding address data", "error", err)
		return domain.Address{}, err
	}

	return responseData.ToAddress(), nil
}
