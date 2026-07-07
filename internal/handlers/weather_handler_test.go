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

	// ACT: CRIAR RECORDER E REQUEST FAKES PARA CAPTURAR E ENVIAR AS CHAMADAS.
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/weather/", nil)
	router.ServeHTTP(recorder, request)

	// ASSERT: VERIFICA RESULTADO
	if recorder.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, recorder.Code)
	}
}
