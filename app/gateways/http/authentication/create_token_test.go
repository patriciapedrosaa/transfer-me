package authentication

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type CreateLoginBadRequest struct {
	cpf int
	secret string
}

func TestLogin(t *testing.T) {
	body := LoginRequest{
		CPF:    "12345678910",
		Secret: "MySecret",
	}
	requestBody, _ := json.Marshal(body)
	responseBody := LoginResponse{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIwMDAxLTAxLTAxVDAwOjE1OjAwWiIsImlhdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiaWQiOiIxODQxZmM1MC1lMjEwLTQyZmYtYmY2YS0wZTIyNjNjYmJiNGUiLCJpc3MiOiJKV1QiLCJuYW1lIjoiUGF0cmljaWEiLCJzdWIiOiI4ZWI3Zjc5Ny1kZWNiLTQ5NTctOGYxOC01NjU4OTA2ODU4MDgifQ.PbwYkUqGlxdb_pjzAEf0dHacc3eaY5C-sCmpEqXzrwo",
	}

	t.Run("should return 201 and generate a Token", func(t *testing.T) {
		handler := createFakeHandler()
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.CreateToken).ServeHTTP(response, request)

		got := response.Body.String()
		expected, _ := json.Marshal(responseBody)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, string(expected), strings.TrimSpace(got))
		assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when body has invalid fields", func(t *testing.T) {
		body := CreateLoginBadRequest{
			cpf:    12345678910,
			secret: "secret",
		}
		requestBody, _ := json.Marshal(body)
		handler := NewHandler(&authentication.UseCaseMock{
			CheckLoginFunc:     nil,
			CreateTokenFunc: func(login authentication.LoginInputs) (string, error) {
				return "", errors.New("invalid fields")
			},
		}, nil)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.CreateToken).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid fields"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when body is empty", func(t *testing.T) {
		body := LoginRequest{
			CPF:    "",
			Secret: "",
		}
		requestBody, _ := json.Marshal(body)
		handler := NewHandler(&authentication.UseCaseMock{
			CheckLoginFunc:     nil,
			CreateTokenFunc: func(login authentication.LoginInputs) (string, error) {
				return "", errors.New("invalid fields")
			},
		}, nil)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.CreateToken).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid fields"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when is missing fields", func(t *testing.T) {
		handler := NewHandler(&authentication.UseCaseMock{
			CheckLoginFunc:     nil,
			CreateTokenFunc: func(login authentication.LoginInputs) (string, error) {
				return "", errors.New("invalid request payload")
			},
		}, nil)
		request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte{}))
		response := httptest.NewRecorder()

		http.HandlerFunc(handler.CreateToken).ServeHTTP(response, request)

		got := response.Body.String()
		expected := `{"error":"invalid request payload"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
	})
}

func createFakeHandler() Handler {
	return NewHandler(&authentication.UseCaseMock{
		CheckLoginFunc: func(inputs authentication.LoginInputs) (bool, error) {
			return true, nil
		},
		CreateTokenFunc: func(login authentication.LoginInputs) (string, error) {
			return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIwMDAxLTAxLTAxVDAwOjE1OjAwWiIsImlhdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiaWQiOiIxODQxZmM1MC1lMjEwLTQyZmYtYmY2YS0wZTIyNjNjYmJiNGUiLCJpc3MiOiJKV1QiLCJuYW1lIjoiUGF0cmljaWEiLCJzdWIiOiI4ZWI3Zjc5Ny1kZWNiLTQ5NTctOGYxOC01NjU4OTA2ODU4MDgifQ.PbwYkUqGlxdb_pjzAEf0dHacc3eaY5C-sCmpEqXzrwo", nil
		},
	}, &account.UseCaseMock{GetByCpfFunc: func(cpf string) (entities.Account, error) {
		return entities.Account{
			CPF:       "12345678910",
			Secret:    "foobar",
			CreatedAt: time.Time{},
		}, nil
	}})
}
