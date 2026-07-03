package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"example.com/address-weather-project/internal/domain"
)

type GeocodingClient struct {
	httpClient *http.Client
	baseUrl    string
	logger     *slog.Logger
}

func NewGeocodingClient(httpClient *http.Client, baseUrl string, logger *slog.Logger) *GeocodingClient {
	return &GeocodingClient{httpClient: httpClient, baseUrl: baseUrl, logger: logger}
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

	c.logger.Info("fetching geocoding", "address", address)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return domain.Localization{}, err
	}

	defer resp.Body.Close()

	c.logger.Info("fetch geocoding status", "postal_code", address.PostalCode, "status_code", resp.StatusCode, "provider", "nominatim")

	if resp.StatusCode != http.StatusOK {
		return domain.Localization{}, fmt.Errorf("nominatim returned non-ok status %d", resp.StatusCode)
	}

	var responseData []NominatimResponse

	err = json.NewDecoder(resp.Body).Decode(&responseData)

	if err != nil {
		c.logger.Error("error decoding geocoding data", "error", err)
		return domain.Localization{}, err
	}

	if len(responseData) == 0 {
		return domain.Localization{}, errors.New("geocoding returned no results")
	}

	return domain.Localization{
		Latitude:  responseData[0].Lat,
		Longitude: responseData[0].Lon,
	}, nil
}
