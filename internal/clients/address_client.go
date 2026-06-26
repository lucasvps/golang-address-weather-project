package clients

import (
	"encoding/json"
	"errors"
	"net/http"

	"example.com/address-weather-project/internal/domain"
)

type AddressClient struct {
	httpClient *http.Client
	baseUrl    string
}

func NewAddressClient(httpClient *http.Client, baseUrl string) *AddressClient {
	return &AddressClient{
		httpClient: httpClient,
		baseUrl:    baseUrl,
	}
}

func (c *AddressClient) FetchAddress(postalCode string) (domain.Address, error) {
	requestUrl := c.baseUrl + postalCode + "/json"

	resp, err := c.httpClient.Get(requestUrl)

	if err != nil {
		return domain.Address{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Address{}, errors.New("viacep returned non-ok status")
	}

	var responseData ViaCepResponse

	err = json.NewDecoder(resp.Body).Decode(&responseData)

	if err != nil {
		return domain.Address{}, err
	}

	return responseData.ToAddress(), nil
}
