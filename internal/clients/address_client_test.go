package clients

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchAddress(t *testing.T) {
	// ARRANGE:
	postalCode := "85035000"
	street := "Rua São Paulo"
	city := "Guarapuava"
	uf := "PR"

	correctBody, _ := json.Marshal(map[string]string{
		"cep":        postalCode,
		"logradouro": street,
		"localidade": city,
		"uf":         uf,
	})

	tableDrivenTests := []struct {
		name          string
		statusCode    int
		responseBody  string
		expectedError bool
		expectedCity  string
	}{
		{
			name:          "success",
			statusCode:    http.StatusOK,
			responseBody:  string(correctBody),
			expectedError: false,
			expectedCity:  city,
		},
		{
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
			//ARRANGE
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(testCase.statusCode)
				w.Write([]byte(testCase.responseBody))
			}))

			defer server.Close()
			addressClient := NewAddressClient(server.Client(), server.URL+"/", slog.Default())

			//ACT
			address, err := addressClient.FetchAddress(postalCode)

			//ASSERT
			if err == nil && testCase.expectedError {
				t.Errorf("expected error, got nil")
			}

			if !testCase.expectedError && err != nil {
				t.Errorf("expected nil error, got %v", err)
			}

			if !testCase.expectedError && err == nil && address.City != testCase.expectedCity {
				t.Errorf("expected city %s, got %s", testCase.expectedCity, address.City)
			}
		})
	}

}
