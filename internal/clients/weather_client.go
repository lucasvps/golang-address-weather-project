package clients

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"example.com/address-weather-project/internal/domain"
)

type WeatherClient struct {
	httpClient *http.Client
	baseUrl    string
	apiKey     string
}

func NewWeatherClient(httpClient *http.Client, baseUrl string, apiKey string) *WeatherClient {
	return &WeatherClient{httpClient: httpClient, baseUrl: baseUrl, apiKey: apiKey}
}

func (c *WeatherClient) FetchWeather(lat string, long string) (domain.Weather, error) {
	params := url.Values{}

	params.Add("lat", lat)
	params.Add("lon", long)
	params.Add("appid", c.apiKey)
	params.Add("units", "metric")

	requestUrl := c.baseUrl + params.Encode()

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)

	if err != nil {
		return domain.Weather{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return domain.Weather{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Weather{}, errors.New("Could not reach the weather data")
	}

	var responseData OpenWeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&responseData)

	if err != nil {
		return domain.Weather{}, err
	}

	return *responseData.ToWeatherDomain(), nil
}
