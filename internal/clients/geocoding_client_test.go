package clients

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/address-weather-project/internal/domain"
)

func TestFetchGeocoding(t *testing.T) {
	latitude := "-25.3914235"
	longitude := "-51.4855767"

	jsonResponseExample := `
		[
			{
				"lat": "-25.3914235",
				"lon": "-51.4855767"
			}
		]`

	tableDrivenTests := []struct {
		name          string
		statusCode    int
		responseBody  string
		expectedError bool
		expectedLat   string
		expectedLon   string
	}{
		{
			name:          "success call",
			statusCode:    http.StatusOK,
			responseBody:  jsonResponseExample,
			expectedError: false,
			expectedLat:   latitude,
			expectedLon:   longitude,
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

	city := "Guarapuava"
	postalCode := "85035000"
	street := "Rua São Paulo"
	format := "json"

	for _, testCase := range tableDrivenTests {
		t.Run(testCase.name, func(t *testing.T) {
			//ARRANGE

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				query := r.URL.Query()

				if query.Get("city") != city {
					t.Errorf("expected city %s, got %s", city, query.Get("city"))
				}

				if query.Get("postalcode") != postalCode {
					t.Errorf("expected postalCode %s, got %s", postalCode, query.Get("postalcode"))
				}

				if query.Get("street") != street {
					t.Errorf("expected street %s, got %s", street, query.Get("street"))
				}

				if query.Get("format") != format {
					t.Errorf("expected format %s, got %s", format, query.Get("format"))
				}

				w.WriteHeader(testCase.statusCode)
				w.Write([]byte(testCase.responseBody))
			}))

			defer server.Close()

			gClient := NewGeocodingClient(server.Client(), server.URL, slog.Default())

			//ACT
			resp, err := gClient.FetchGeocoding(domain.Address{
				City:       city,
				PostalCode: postalCode,
				Street:     street,
			})

			//ASSERT
			if testCase.expectedError && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !testCase.expectedError && err != nil {
				t.Errorf("expected nil error, got %v", err)
			}

			if !testCase.expectedError && testCase.expectedLat != resp.Latitude {
				t.Errorf("expected latitude %v but got %v", testCase.expectedLat, resp.Latitude)
			}

			if !testCase.expectedError && testCase.expectedLon != resp.Longitude {
				t.Errorf("expected longitude %v but got %v", testCase.expectedLon, resp.Longitude)
			}

		})
	}

}
