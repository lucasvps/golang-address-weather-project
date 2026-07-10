package clients

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchWeather(t *testing.T) {
	// ARRANGE

	temperature := 29.86

	jsonResponseExample := `{
    "coord": {
        "lon": 10.99,
        "lat": 44.34
    },
    "weather": [
        {
            "id": 500,
            "main": "Rain",
            "description": "light rain",
            "icon": "10d"
        }
    ],
    "base": "stations",
    "main": {
        "temp": 29.86,
        "feels_like": 32.76,
        "temp_min": 29.56,
        "temp_max": 30.7,
        "pressure": 1011,
        "humidity": 61,
        "sea_level": 1011,
        "grnd_level": 946
    },
    "visibility": 10000,
    "wind": {
        "speed": 1.94,
        "deg": 253,
        "gust": 4.14
    },
    "rain": {
        "1h": 0.43
    },
    "clouds": {
        "all": 26
    },
    "dt": 1782906768,
    "sys": {
        "type": 2,
        "id": 2004688,
        "country": "IT",
        "sunrise": 1782876930,
        "sunset": 1782932635
    },
    "timezone": 7200,
    "id": 3163858,
    "name": "Zocca",
    "cod": 200
}`

	tableDrivenTests := []struct {
		name                string
		statusCode          int
		responseBody        string
		expectedError       bool
		expectedTemperature float64
	}{
		{
			name:                "success call",
			statusCode:          http.StatusOK,
			responseBody:        jsonResponseExample,
			expectedError:       false,
			expectedTemperature: temperature,
		}, {
			name:          "non ok status",
			statusCode:    http.StatusInternalServerError,
			responseBody:  `{}`,
			expectedError: true,
		},
		{
			name:          "invalid json",
			statusCode:    http.StatusOK,
			responseBody:  `not-json`,
			expectedError: true,
		},
	}

	for _, testCase := range tableDrivenTests {
		t.Run(testCase.name, func(t *testing.T) {
			// ARRANGE
			lat := "832918321"
			long := "-382193821"

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				query := r.URL.Query()

				if query.Get("lat") != lat {
					t.Errorf("expected lat %s, got %s", lat, query.Get("lat"))
				}

				if query.Get("lon") != long {
					t.Errorf("expected lon %s, got %s", long, query.Get("lon"))
				}

				if query.Get("appid") != "api-key" {
					t.Errorf("expected appid %s, got %s", "api-key", query.Get("appid"))
				}

				if query.Get("units") != "metric" {
					t.Errorf("expected units %s, got %s", "metric", query.Get("units"))
				}

				w.WriteHeader(testCase.statusCode)
				w.Write([]byte(testCase.responseBody))
			}))

			defer server.Close()

			wClient := NewWeatherClient(server.Client(), server.URL, "api-key", slog.Default())

			//ACT
			resp, err := wClient.FetchWeather(lat, long)

			//ASSERT
			if testCase.expectedError && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !testCase.expectedError && err != nil {
				t.Errorf("expected nil error, got %v", err)
			}

			if !testCase.expectedError && testCase.expectedTemperature != resp.Temperature {
				t.Errorf("expected temperature %v but got %v", testCase.expectedTemperature, resp.Temperature)
			}
		})
	}

}
