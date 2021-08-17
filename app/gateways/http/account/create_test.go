package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type CreateAccountBadRequest struct {
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Secret string `json:"secret"`
}

func TestCreate(t *testing.T) {
	body := CreateAccountRequest{
		Name:   "Peter Park",
		CPF:    "12345678910",
		Secret: "MySecret",
	}
	requestBody, _ := json.Marshal(body)
	responseBody := CreateAccountResponse{
		AccountID: "6a00ac20-e07f-455f-a53c-37088c7b4266",
		Name:      "Peter Park",
	}
	t.Run("should return 201 and created account", func(t *testing.T) {
		handler := createFakeHandler(nil)
		request, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Create).ServeHTTP(response, request)

		got := response.Body.String()
		expected, _ := json.Marshal(responseBody)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, string(expected), strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))

	})

	t.Run("should return 400 and error when body has invalid fields", func(t *testing.T) {
		body := CreateAccountRequest{
			Name:   "Peter Park",
			CPF:    "one two three",
			Secret: "MySecret",
		}
		requestBody, _ := json.Marshal(body)
		err := errors.New("invalid cpf")
		handler := createFakeHandler(err)
		request, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Create).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid cpf"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, got)
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when body has invalid type fields", func(t *testing.T) {
		body := CreateAccountBadRequest{
			Name:   "Peter Park",
			CPF:    12345678910,
			Secret: "MySecret",
		}
		requestBody, _ := json.Marshal(body)
		handler := NewHandler(nil)
		request, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Create).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid request payload"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, got)
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when body is empty", func(t *testing.T) {
		body := CreateAccountRequest{
			Name:   "",
			CPF:    "",
			Secret: "",
		}
		requestBody, _ := json.Marshal(body)
		handler := createFakeHandler(nil)
		request, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Create).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid fields"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, got)
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when is missing fields", func(t *testing.T) {
		handler := createFakeHandler(nil)
		request, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader([]byte{}))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Create).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid request payload"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, got)
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

}

func createFakeHandler(err error) Handler {
	if err != nil {
		return NewHandler(&account.UseCaseMock{
			CreateFunc: func(input account.CreateAccountInput) (entities.Account, error) {
				return entities.Account{}, err
			},
		})
	}
	return NewHandler(&account.UseCaseMock{
		CreateFunc: func(input account.CreateAccountInput) (entities.Account, error) {
			return entities.Account{
				AccountID: "6a00ac20-e07f-455f-a53c-37088c7b4266",
				Name:      input.Name,
				CPF:       vos.CPF(input.CPF),
				Secret:    "hash",
				Balance:   100,
				CreatedAt: time.Now(),
			}, nil
		},
	})
}
