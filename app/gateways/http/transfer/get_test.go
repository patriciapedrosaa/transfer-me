package transfer

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var timeCreation = time.Now()

var getHeader = http.Header{
	"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjk0ODI3MzQsImlhdCI6MTYyOTQ4MTgzNCwiaWQiOiJkY2E3NzQ2ZC04YWU1LTQ3Y2UtOGExYi0yOGFhOTFhZjkyNWQiLCJpc3MiOiJKV1QiLCJuYW1lIjoiUGF0cmljaWEiLCJzdWIiOiI2NDJlMGQ0NC05NzkyLTRkNmYtOWEwNC1iNDAxODZkZGRiZWYifQ.xqFGOp_3jatWFPLAxe9WtvRSITV1FgQzPAnePwXA2EE"},
	"Content-Type":  []string{"application/json"},
}

func TestGet(t *testing.T) {
	t.Run("should return 200 and a list of transfers", func(t *testing.T) {
		responseBody := createFakeGetTransferResponse()
		handler := createGetFakeHandler(nil)
		request, _ := http.NewRequest(http.MethodGet, "/transfers", nil)
		request.Header = getHeader
		response := httptest.NewRecorder()
		ctx := context.WithValue(request.Context(), http_server.AccountID, "642e0d44-9792-4d6f-9a04-b40186dddbef")

		http.HandlerFunc(handler.Get).ServeHTTP(response, request.WithContext(ctx))

		got := response.Body.String()
		expected, _ := json.Marshal(responseBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, string(expected), got)
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

	t.Run("should return 500 and an error message", func(t *testing.T) {
		err := errors.New(ErrUnexpected)
		handler := createGetFakeHandler(err)
		request, _ := http.NewRequest(http.MethodGet, "/transfers", nil)
		request.Header = header
		response := httptest.NewRecorder()
		ctx := context.WithValue(request.Context(), http_server.AccountID, "642e0d44-9792-4d6f-9a04-b40186dddbef")

		http.HandlerFunc(handler.Get).ServeHTTP(response, request.WithContext(ctx))

		got := response.Body.String()
		expected := `{"error":"something went wrong"}`

		assert.Equal(t, http.StatusForbidden, response.Code)
		assert.Equal(t, expected, strings.TrimSpace(got))
		assert.Equal(t, http_server.JsonContentType, response.Header().Get("Content-Type"))
	})

}

func createGetFakeHandler(err error) Handler {
	if err != nil {
		return NewHandler(&transfer.UseCaseMock{
			GetTransfersByAccountIDFunc: func(accountID string) ([]entities.Transfer, error) {
				return nil, err
			},
		}, nil)
	}
	return NewHandler(&transfer.UseCaseMock{
		GetTransfersByAccountIDFunc: func(accountID string) ([]entities.Transfer, error) {
			return []entities.Transfer{
				{
					TransferID:           "6a00ac20-e07f-455f-a53c-37088c7b4266",
					AccountOriginID:      "642e0d44-9792-4d6f-9a04-b40186dddbef",
					AccountDestinationID: "6a00ac20-e07f-455f-a53c-37088c7b4277",
					Amount:               50,
					CreatedAt:            timeCreation,
				},
				{
					TransferID:           "6a00ac20-e07f-455f-a53c-37088c7b4267",
					AccountOriginID:      "642e0d44-9792-4d6f-9a04-b40186dddbef",
					AccountDestinationID: "6a00ac20-e07f-455f-a53c-37088c7b4277",
					Amount:               50,
					CreatedAt:            timeCreation,
				},
			}, nil
		},
	}, nil)
}

func createFakeGetTransferResponse() []GetTransferResponse {
	return []GetTransferResponse{
		{
			TransferID: "6a00ac20-e07f-455f-a53c-37088c7b4266",
			CreatedAt:  timeCreation,
		},
		{
			TransferID: "6a00ac20-e07f-455f-a53c-37088c7b4267",
			CreatedAt:  timeCreation,
		},
	}
}
