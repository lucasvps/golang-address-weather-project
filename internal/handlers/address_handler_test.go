package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/address-weather-project/internal/domain"
	"github.com/gin-gonic/gin"
)

//MOCKS

type MockAddressService struct {
	response *domain.Address
	err      error
}

func (m MockAddressService) FetchAddress(postalCode string) (*domain.Address, error) {
	return m.response, m.err
}

func TestFetchAddressFromPostalCode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// ARRANGE: CRIA MOCK DO SERVICE E INSTANCIA HANDLER
	mockAddressService := &MockAddressService{}
	handler := NewAddressHandler(mockAddressService)
	router := gin.New()
	router.GET("/address/:postalCode", handler.FetchAddressFromPostalCode)

	// ACT
	tableDrivenTests := []struct {
		name           string
		path           string
		expectedStatus int
	}{
		{
			name:           "invalid postal code bigger",
			path:           "/address/38219382198",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid postal code smaller",
			path:           "/address/12345",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid postal code letters",
			path:           "/address/8503K000",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "missing postal code",
			path:           "/address/",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "valid postal code fetch",
			path:           "/address/85035000",
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

	// ASSERT
}
