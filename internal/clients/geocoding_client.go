package clients

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"example.com/address-weather-project/internal/domain"
)

type GeocodingClient struct {
	httpClient *http.Client
	baseUrl    string
}

func NewGeocodingClient(httpClient *http.Client, baseUrl string) *GeocodingClient {
	return &GeocodingClient{httpClient: httpClient, baseUrl: baseUrl}
}

func (c *GeocodingClient) FetchGeocoding(address domain.Address) (domain.Localization, error) {
	params := url.Values{}

	params.Add("city", address.City)
	params.Add("postalcode", address.PostalCode)
	params.Add("street", address.Street)
	params.Add("country", "Brazil")
	params.Add("format", "json")
	params.Add("limit", "1")

	requestUrl := c.baseUrl + "?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)

	if err != nil {
		return domain.Localization{}, err
	}

	req.Header.Add("User-Agent", "address-weather-project/1.0")

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return domain.Localization{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Localization{}, errors.New("An error occurred while fetching the localization.")
	}

	var responseData []NominatimResponse

	err = json.NewDecoder(resp.Body).Decode(&responseData)

	if err != nil {
		return domain.Localization{}, err
	}

	if len(responseData) == 0 {
		return domain.Localization{}, errors.New("We cannot found the localization.")
	}

	return domain.Localization{
		Latitude:  responseData[0].Lat,
		Longitude: responseData[0].Lon,
	}, nil
}
