package account

import (
	"encoding/json"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetBalance(t *testing.T) {
	t.Run("should return 200 and balance", func(t *testing.T) {
		responseBody := GetBalanceResponse{Balance: 100}
		handler := NewHandler(&account.UseCaseMock{GetBalanceFunc: func(id string) (int, error) {
			return 100, nil
		}}, zerolog.Logger{})
		request, _ := http.NewRequest(http.MethodGet, "/accounts/id/balance", nil)
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.GetBalance).ServeHTTP(response, request)

		got := response.Body.String()
		expected, _ := json.Marshal(responseBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, got, string(expected))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 404 when account was not found", func(t *testing.T) {
		handler := NewHandler(&account.UseCaseMock{GetBalanceFunc: func(id string) (int, error) {
			return 0, errNotFound
		}}, zerolog.Logger{})
		request, _ := http.NewRequest(http.MethodGet, "/accounts/id/balance", nil)
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.GetBalance).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"not found"}`

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, got, expected)
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 500 and internal server error", func(t *testing.T) {
		handler := NewHandler(&account.UseCaseMock{GetBalanceFunc: func(id string) (int, error) {
			return 0, errors.New("something went wrong")
		}}, zerolog.Logger{})
		request, _ := http.NewRequest(http.MethodGet, "/accounts/id/balance", nil)
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.GetBalance).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"something went wrong"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, got, expected)
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})
}
