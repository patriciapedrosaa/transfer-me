package authentication

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication/usecase"
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

const TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIwMDAxLTAxLTAxVDAwOjE1OjAwWiIsImlhdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiaWQiOiIxODQxZmM1MC1lMjEwLTQyZmYtYmY2YS0wZTIyNjNjYmJiNGUiLCJpc3MiOiJKV1QiLCJuYW1lIjoiUGF0cmljaWEiLCJzdWIiOiI4ZWI3Zjc5Ny1kZWNiLTQ5NTctOGYxOC01NjU4OTA2ODU4MDgifQ.PbwYkUqGlxdb_pjzAEf0dHacc3eaY5C-sCmpEqXzrwo"

type LoginBadRequest struct {
	cpf    int
	secret string
}

func TestLogin(t *testing.T) {
	t.Run("should return 201 and generate a Token", func(t *testing.T) {
		body := LoginRequest{
			CPF:    "12345678910",
			Secret: "MySecret",
		}
		requestBody, _ := json.Marshal(body)
		responseBody := LoginResponse{
			Token: TOKEN,
		}
		handler := createFakeHandler(TOKEN, nil, nil)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		got := response.Body.String()
		expected, _ := json.Marshal(responseBody)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, string(expected), strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when body has invalid fields", func(t *testing.T) {
		body := LoginBadRequest{
			cpf:    12345678910,
			secret: "secret",
		}
		requestBody, _ := json.Marshal(body)
		handler := createFakeHandler("", nil, nil)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid fields"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when body is empty", func(t *testing.T) {
		body := LoginRequest{
			CPF:    "",
			Secret: "",
		}
		requestBody, _ := json.Marshal(body)
		handler := createFakeHandler("", nil, nil)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid fields"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when is missing fields", func(t *testing.T) {
		handler := createFakeHandler("", nil, nil)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte{}))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid request payload"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when CPF does not exist", func(t *testing.T) {
		body := LoginRequest{
			CPF:    "12345678913",
			Secret: "MySecret",
		}
		requestBody, _ := json.Marshal(body)
		err := usecase.ErrInvalidCPF
		handler := createFakeHandler("", err, nil)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"incorrect username or password"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when secret is incorrect", func(t *testing.T) {
		body := LoginRequest{
			CPF:    "12345678910",
			Secret: "MySecret",
		}
		requestBody, _ := json.Marshal(body)
		err := usecase.ErrInvalidSecret
		handler := createFakeHandler("", nil, err)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"incorrect username or password"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 500 and error when some unexpected error occurs", func(t *testing.T) {
		body := LoginRequest{
			CPF:    "12345678910",
			Secret: "MySecret",
		}
		requestBody, _ := json.Marshal(body)
		err := errors.New("unexpected error")
		handler := createFakeHandler("", nil, err)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Login).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"unexpected error"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})
}

func createFakeHandler(token string, errGetByCPF error, errCreateToken error) Handler {
	return NewHandler(&authentication.UseCaseMock{
		CreateTokenFunc: func(ctx context.Context, login authentication.LoginInputs) (string, error) {
			if errCreateToken != nil {
				return "", errCreateToken
			}
			return token, nil
		},
	}, &account.UseCaseMock{
		GetByCpfFunc: func(ctx context.Context, cpf string) (entities.Account, error) {
			if errGetByCPF != nil {
				return entities.Account{}, errGetByCPF
			}
			return entities.Account{
				AccountID: "1d773d5d-2f8d-4e7c-99cb-8b62b8d07d41",
				Name:      "Peter Park",
				CPF:       "12345678910",
				Secret:    "foobar",
				Balance:   100,
				CreatedAt: time.Now(),
			}, nil
		}}, zerolog.Logger{})
}
