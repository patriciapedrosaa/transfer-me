package account

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	t.Run("should return 200 and a empty list of accounts", func(t *testing.T) {
		handler := NewHandler(&account.UseCaseMock{
			GetAccountsFunc: func(ctx context.Context) ([]entities.Account, error) {
				return []entities.Account{}, nil
			},
		}, zerolog.Logger{})

		request, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Get).ServeHTTP(response, request)

		got := response.Body.String()
		expected := "[]"

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expected, got)
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 200 and a list of accounts", func(t *testing.T) {
		responseBody := generateFakeGetAccountResponse()
		handler := generateFakeGetHandler()
		request, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Get).ServeHTTP(response, request)

		got := response.Body.String()
		expected, _ := json.Marshal(responseBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, string(expected), got)
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 500 and a error message", func(t *testing.T) {
		handler := NewHandler(&account.UseCaseMock{
			GetAccountsFunc: func(ctx context.Context) ([]entities.Account, error) {
				return nil, errors.New("something went wrong")
			}}, zerolog.Logger{})
		request, _ := http.NewRequest(http.MethodGet, "/accounts", nil)
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Get).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"something went wrong"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

}

func generateFakeGetHandler() Handler {
	return NewHandler(&account.UseCaseMock{
		GetAccountsFunc: func(ctx context.Context) ([]entities.Account, error) {
			return []entities.Account{
				{
					AccountID: "6a00ac20-e07f-455f-a53c-37088c7b4266",
					Name:      "Jack Sparrow",
					CPF:       "12345678910",
					Secret:    "hash",
					Balance:   100,
					CreatedAt: time.Now(),
				},
				{
					AccountID: "6a00ac20-e07f-455f-a53c-37088c7b4267",
					Name:      "William Turner",
					CPF:       "12345678911",
					Secret:    "hash",
					Balance:   100,
					CreatedAt: time.Now(),
				},
			}, nil
		},
	}, zerolog.Logger{})
}

func generateFakeGetAccountResponse() []GetAccountResponse {
	return []GetAccountResponse{
		{
			AccountID: "6a00ac20-e07f-455f-a53c-37088c7b4266",
			Name:      "Jack Sparrow",
			CPF:       "12345678910",
		},
		{
			AccountID: "6a00ac20-e07f-455f-a53c-37088c7b4267",
			Name:      "William Turner",
			CPF:       "12345678911",
		},
	}
}
