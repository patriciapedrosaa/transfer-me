package transfer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer/usecase"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var header = http.Header{
	"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjk0ODI3MzQsImlhdCI6MTYyOTQ4MTgzNCwiaWQiOiJkY2E3NzQ2ZC04YWU1LTQ3Y2UtOGExYi0yOGFhOTFhZjkyNWQiLCJpc3MiOiJKV1QiLCJuYW1lIjoiUGF0cmljaWEiLCJzdWIiOiI2NDJlMGQ0NC05NzkyLTRkNmYtOWEwNC1iNDAxODZkZGRiZWYifQ.xqFGOp_3jatWFPLAxe9WtvRSITV1FgQzPAnePwXA2EE"},
	"Content-Type":  []string{"application/json"},
}

type CreateTransferBadRequest struct {
	DestinationAccountID string `json:"account_destination_id" validate:"required"`
	Amount               string `json:"amount" validate:"required"`
}

func TestCreate(t *testing.T) {
	body := CreateTransferRequest{
		DestinationAccountID: "6a00ac20-e07f-455f-a53c-37088c7b4266",
		Amount:               50,
	}
	requestBody, _ := json.Marshal(body)
	responseBody := CreateTransferResponse{
		TransferID: "0de9ec06-0ca4-4583-9ddc-585ec65a8c29",
	}

	t.Run("should return 201 and created transfer", func(t *testing.T) {
		transferID := "0de9ec06-0ca4-4583-9ddc-585ec65a8c29"
		accountOriginID := "642e0d44-9792-4d6f-9a04-b40186dddbef"
		accountDestinationID := "6a00ac20-e07f-455f-a53c-37088c7b4266"
		amount := 50

		handler := createFakeHandler(transferID, accountOriginID, accountDestinationID, amount, nil)
		request, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(requestBody))
		ctx := context.WithValue(request.Context(), http_server.AccountID, "642e0d44-9792-4d6f-9a04-b40186dddbef")

		response := httptest.NewRecorder()
		request.Header = header

		handler.Create(response, request.WithContext(ctx))

		got := response.Body.String()
		expected, _ := json.Marshal(responseBody)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, string(expected), strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when body has invalid field types", func(t *testing.T) {
		badBody := CreateTransferBadRequest{
			DestinationAccountID: "6a00ac20-e07f-455f-a53c-37088c7b4266",
			Amount:               "fifty",
		}
		badRequestBody, _ := json.Marshal(badBody)
		handler := createFakeHandler("", "", "", 0, nil)
		request, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(badRequestBody))
		ctx := context.WithValue(request.Context(), http_server.AccountID, "642e0d44-9792-4d6f-9a04-b40186dddbef")

		response := httptest.NewRecorder()
		request.Header = header

		handler.Create(response, request.WithContext(ctx))

		got := response.Body.String()
		expected := `{"error":"invalid request payload"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when body is empty", func(t *testing.T) {
		emptyBody := CreateTransferRequest{
			DestinationAccountID: "",
			Amount:               0,
		}
		emptyRequestBody, _ := json.Marshal(emptyBody)
		handler := createFakeHandler("", "", "", 0, nil)
		request, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(emptyRequestBody))
		ctx := context.WithValue(request.Context(), http_server.AccountID, "642e0d44-9792-4d6f-9a04-b40186dddbef")

		response := httptest.NewRecorder()
		request.Header = header

		handler.Create(response, request.WithContext(ctx))

		got := response.Body.String()
		expected := `{"error":"invalid fields"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when is missing fields", func(t *testing.T) {
		handler := createFakeHandler("", "", "", 0, nil)
		request, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader([]byte{}))
		ctx := context.WithValue(request.Context(), http_server.AccountID, "642e0d44-9792-4d6f-9a04-b40186dddbef")
		response := httptest.NewRecorder()
		request.Header = header
		handler.Create(response, request.WithContext(ctx))

		got := response.Body.String()
		expected := `{"error":"invalid request payload"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 400 and error when cpf is not found", func(t *testing.T) {
		transferID := "0de9ec06-0ca4-4583-9ddc-585ec65a8c29"
		accountOriginID := "642e0d44-9792-4d6f-9a04-b40186dddbef"
		accountDestinationID := "6a00ac20-e07f-455f-a53c-37088c7b4255"
		amount := 50
		err := usecase.ErrNotFound
		handler := createFakeHandler(transferID, accountOriginID, accountDestinationID, amount, err)
		request, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(requestBody))
		ctx := context.WithValue(request.Context(), http_server.ContextID, "642e0d44-9792-4d6f-9a04-b40186dddbef")
		response := httptest.NewRecorder()
		request.Header = header
		handler.Create(response, request.WithContext(ctx))

		got := response.Body.String()
		expected := `{"error":"invalid transfer data"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 500 and error when an unexpected error occurs", func(t *testing.T) {
		transferID := "0de9ec06-0ca4-4583-9ddc-585ec65a8c29"
		accountOriginID := "642e0d44-9792-4d6f-9a04-b40186dddbef"
		accountDestinationID := "6a00ac20-e07f-455f-a53c-37088c7b4255"
		amount := 50
		err := usecase.ErrUnexpected
		handler := createFakeHandler(transferID, accountOriginID, accountDestinationID, amount, err)
		request, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(requestBody))
		ctx := context.WithValue(request.Context(), http_server.AccountID, "642e0d44-9792-4d6f-9a04-b40186dddbef")
		response := httptest.NewRecorder()
		request.Header = header
		handler.Create(response, request.WithContext(ctx))

		got := response.Body.String()
		expected := `{"error":"something went wrong"}`

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 403 and error when amount is invalid", func(t *testing.T) {
		transferID := "0de9ec06-0ca4-4583-9ddc-585ec65a8c29"
		accountOriginID := "642e0d44-9792-4d6f-9a04-b40186dddbef"
		accountDestinationID := "6a00ac20-e07f-455f-a53c-37088c7b4255"
		amount := -50
		err := errors.New("the amount must be greater than zero")
		handler := createFakeHandler(transferID, accountOriginID, accountDestinationID, amount, err)
		request, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(requestBody))
		ctx := context.WithValue(request.Context(), http_server.AccountID, "642e0d44-9792-4d6f-9a04-b40186dddbef")
		response := httptest.NewRecorder()
		request.Header = header
		handler.Create(response, request.WithContext(ctx))

		got := response.Body.String()
		expected := `{"error":"the amount must be greater than zero"}`

		assert.Equal(t, http.StatusForbidden, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

}

func createFakeHandler(transferID, AccountOriginID, AccountDestinationID string, amount int, err error) Handler {
	if err != nil {
		return NewHandler(&transfer.UseCaseMock{
			CreateFunc: func(input transfer.CreateTransferInput) (entities.Transfer, error) {
				return entities.Transfer{}, err
			},
		}, nil, nil)
	}
	return NewHandler(&transfer.UseCaseMock{
		CreateFunc: func(input transfer.CreateTransferInput) (entities.Transfer, error) {
			return entities.Transfer{
				TransferID:           transferID,
				AccountOriginID:      AccountOriginID,
				AccountDestinationID: AccountDestinationID,
				Amount:               amount,
				CreatedAt:            time.Now(),
			}, nil
		},
	}, &authentication.UseCaseMock{
		ValidatesTokenFunc: func(tokenString string) (entities.Token, error) {
			return entities.Token{
				ID:        "6a00ac20-e07f-455f-a53c-37088c7b4266",
				Name:      "Olive Oyl",
				Subject:   "642e0d44-9792-4d6f-9a04-b40186dddbef",
				Issuer:    "JWT",
				IssuedAt:  time.Now(),
				ExpiredAt: time.Now().Add(time.Minute * 15),
			}, nil
		},
	}, &account.UseCaseMock{
		UpdateBalanceFunc: func(originAccountId string, destinationAccountId string, amount int) error {
			return nil
		},
	})
}
