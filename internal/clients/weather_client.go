package clients

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"example.com/address-weather-project/internal/domain"
)

type WeatherClient struct {
	httpClient *http.Client
	baseUrl    string
	apiKey     string
	logger     *slog.Logger
}

func NewWeatherClient(httpClient *http.Client, baseUrl string, apiKey string, logger *slog.Logger) *WeatherClient {
	return &WeatherClient{httpClient: httpClient, baseUrl: baseUrl, apiKey: apiKey, logger: logger}
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

	c.logger.Info("fetching weather", "lat", lat, "lon", long)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return domain.Weather{}, err
	}

	defer resp.Body.Close()

	c.logger.Info("fetch weather status", "status_code", resp.StatusCode, "provider", "openweather")

	if resp.StatusCode != http.StatusOK {
		return domain.Weather{}, fmt.Errorf("openweather returned non-ok status %d", resp.StatusCode)
	}

	var responseData OpenWeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&responseData)

	if err != nil {
		c.logger.Error("error decoding weather data", "error", err)
		return domain.Weather{}, err
	}

	return *responseData.ToWeatherDomain(), nil
}
