package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/address-weather-project/internal/domain"
	"github.com/gin-gonic/gin"
)

// MOCKS MANUAL
type MockWeatherService struct {
	response *domain.WeatherResponse
	err      error
}

func (m MockWeatherService) FetchWeatherDataFromPostalCode(postalCode string) (*domain.WeatherResponse, error) {
	return m.response, m.err
}

// TEST
func TestFetchWeatherByPostalCode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// ARRANGE: CRIA MOCK DO SERVICE E INSTANCIA HANDLER
	mockWeatherService := &MockWeatherService{}
	handler := NewWeatherHandler(mockWeatherService)
	router := gin.New()
	router.GET("/weather/:postalCode", handler.FetchWeatherByPostalCode)

	tableDrivenTests := []struct {
		name           string
		path           string
		expectedStatus int
	}{
		{
			name:           "invalid postal code bigger than 9 char",
			path:           "/weather/85035321893218",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid postal code smaller than 9 char",
			path:           "/weather/8492",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "missing postal code",
			path:           "/weather/",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "valid postal code",
			path:           "/weather/85035000",
			expectedStatus: http.StatusOK,
		},
	}

	for _, testCase := range tableDrivenTests {
		t.Run(testCase.name, func(t *testing.T) {
			// ACT: CRIAR RECORDER E REQUEST FAKES PARA CAPTURAR E ENVIAR AS CHAMADAS.
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodGet, testCase.path, nil)
			router.ServeHTTP(recorder, request)

			// ASSERT: VERIFICA RESULTADO
			if recorder.Code != testCase.expectedStatus {
				t.Errorf("expected status %d, got %d", testCase.expectedStatus, recorder.Code)
			}
		})
	}
}
