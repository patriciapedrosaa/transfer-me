package transfer

import (
	"encoding/json"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer/usecase"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"net/http"
)

var (
	ErrInvalidPayload      = "invalid request payload"
	ErrRequiredFields      = "invalid fields"
	ErrInvalidDataTransfer = "invalid transfer data"
	ErrUnexpected          = "something went wrong"
)

type CreateTransferRequest struct {
	DestinationAccountID string `json:"account_destination_id" validate:"required"`
	Amount               int    `json:"amount" validate:"required"`
}

type CreateTransferResponse struct {
	TransferID string `json:"id" validate:"required"`
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	accountOriginID := r.Context().Value(http_server.AccountID).(string)

	var body CreateTransferRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		http_server.ResponseError(w, http.StatusBadRequest, ErrInvalidPayload)
		return
	}

	validator := http_server.NewJSONValidator()
	err := validator.Validate(body)
	if err != nil {
		http_server.ResponseError(w, http.StatusBadRequest, ErrRequiredFields)
		return
	}

	defer r.Body.Close()

	input := transfer.CreateTransferInput{
		OriginAccountId:      accountOriginID,
		DestinationAccountId: body.DestinationAccountID,
		Amount:               body.Amount,
	}

	output, err := h.useCase.Create(input)

	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrNotFound):
			http_server.ResponseError(w, http.StatusBadRequest, ErrInvalidDataTransfer)
		case errors.Is(err, usecase.ErrUnexpected):
			http_server.ResponseError(w, http.StatusInternalServerError, err.Error())
		default:
			http_server.ResponseError(w, http.StatusForbidden, err.Error())
		}
		return
	}

	err = h.accountUseCase.UpdateBalance(input.OriginAccountId, input.DestinationAccountId, input.Amount)
	if err != nil {
		http_server.ResponseError(w, http.StatusInternalServerError, ErrUnexpected)
		return
	}

	response := CreateTransferResponse{
		TransferID: output.TransferID,
	}

	http_server.ResponseSuccess(w, http.StatusCreated, response)
}
