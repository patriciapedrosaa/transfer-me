package transfer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	account_usecase "github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
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

var fakeAccount = entities.Account{
	AccountID: uuid.New().String(),
	Name:      "Peter Park",
	CPF:       "12345678910",
	Secret:    "secret",
	Balance:   100,
	CreatedAt: time.Now().UTC(),
}

var fakeTransfer = entities.Transfer{
	TransferID:           "0de9ec06-0ca4-4583-9ddc-585ec65a8c29",
	AccountOriginID:      "642e0d44-9792-4d6f-9a04-b40186dddbef",
	AccountDestinationID: "6a00ac20-e07f-455f-a53c-37088c7b4266",
	Amount:               50,
	CreatedAt:            time.Now().UTC(),
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

		handler := createFakeHandler(transferID, accountOriginID, accountDestinationID, amount, nil, nil, nil)
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
		handler := createFakeHandler("", "", "", 0, nil, nil, nil)
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
		handler := createFakeHandler("", "", "", 0, nil, nil, nil)
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
		handler := createFakeHandler("", "", "", 0, nil, nil, nil)
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

	t.Run("should return 400 and error when destination cpf is not found", func(t *testing.T) {
		badBody := CreateTransferRequest{
			DestinationAccountID: "6a00ac20-e07f-455f-a53c-37088c7b4255",
			Amount:               50,
		}
		badRequestBody, _ := json.Marshal(badBody)
		transferID := "0de9ec06-0ca4-4583-9ddc-585ec65a8c29"
		accountOriginID := "642e0d44-9792-4d6f-9a04-b40186dddbef"
		accountDestinationID := badBody.DestinationAccountID
		amount := badBody.Amount
		err := account_usecase.ErrNotFound
		handler := createFakeHandler(transferID, accountOriginID, accountDestinationID, amount, nil, err, nil)
		request, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(badRequestBody))
		ctx := context.WithValue(request.Context(), http_server.AccountID, "642e0d44-9792-4d6f-9a04-b40186dddbef")
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
		handler := createFakeHandler(transferID, accountOriginID, accountDestinationID, amount, err, nil, nil)
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

	t.Run("should return 400 and error when amount is invalid", func(t *testing.T) {
		transferID := "0de9ec06-0ca4-4583-9ddc-585ec65a8c29"
		accountOriginID := "642e0d44-9792-4d6f-9a04-b40186dddbef"
		accountDestinationID := "6a00ac20-e07f-455f-a53c-37088c7b4255"
		amount := -50
		err := errors.New("the amount must be greater than zero")
		handler := createFakeHandler(transferID, accountOriginID, accountDestinationID, amount, err, nil, nil)
		request, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(requestBody))
		ctx := context.WithValue(request.Context(), http_server.AccountID, "642e0d44-9792-4d6f-9a04-b40186dddbef")
		response := httptest.NewRecorder()
		request.Header = header
		handler.Create(response, request.WithContext(ctx))

		got := response.Body.String()
		expected := `{"error":"the amount must be greater than zero"}`

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 500 and error when an error to update balance occurs", func(t *testing.T) {
		transferID := "0de9ec06-0ca4-4583-9ddc-585ec65a8c29"
		accountOriginID := "642e0d44-9792-4d6f-9a04-b40186dddbef"
		accountDestinationID := "6a00ac20-e07f-455f-a53c-37088c7b4266"
		amount := 50
		err := errors.New("can not update balance")
		handler := createFakeHandler(transferID, accountOriginID, accountDestinationID, amount, nil, nil, err)
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

}

func createFakeHandler(transferID, AccountOriginID, AccountDestinationID string, amount int, err error, accountErr error, errUpdateBalance error) Handler {
	if err != nil {
		return NewHandler(&transfer.UseCaseMock{
			CreateFunc: func(input transfer.CreateTransferInput) (entities.Transfer, error) {
				return entities.Transfer{}, err
			},
		}, &account.UseCaseMock{
			GetByIdFunc: func(id string) (entities.Account, error) {
				return entities.Account{}, nil
			},
			UpdateBalanceFunc: func(originAccountId string, destinationAccountId string, amount int) error {
				return nil
			},
		})
	}
	if accountErr != nil {
		return NewHandler(nil, &account.UseCaseMock{
			GetByIdFunc: func(id string) (entities.Account, error) {
				return entities.Account{}, accountErr
			},
		})
	}
	if errUpdateBalance != nil {
		return NewHandler(&transfer.UseCaseMock{
			CreateFunc: func(input transfer.CreateTransferInput) (entities.Transfer, error) {
				return fakeTransfer, nil
			},
		}, &account.UseCaseMock{
			UpdateBalanceFunc: func(originAccountId string, destinationAccountId string, amount int) error {
				return errUpdateBalance
			},
		})
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
	}, &account.UseCaseMock{
		GetByIdFunc: func(id string) (entities.Account, error) {
			return fakeAccount, nil
		},
		UpdateBalanceFunc: func(originAccountId string, destinationAccountId string, amount int) error {
			return nil
		},
	})
}
