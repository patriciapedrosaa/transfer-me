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

	h.logger.Info().Msgf("Create transfer request: %v", body)

	if err := decoder.Decode(&body); err != nil {
		h.logger.Err(err).
			Int("Status_code", http.StatusBadRequest).
			Msg("Occurred when decoding body")
		http_server.ResponseError(w, http.StatusBadRequest, ErrInvalidPayload)
		return
	}

	validator := http_server.NewJSONValidator()
	err := validator.Validate(body)
	if err != nil {
		h.logger.Err(err).
			Int("Status_code", http.StatusBadRequest).
			Msg("Occurred when was validating body")
		http_server.ResponseError(w, http.StatusBadRequest, ErrRequiredFields)
		return
	}

	defer r.Body.Close()

	originAccount, err := h.accountUseCase.GetById(accountOriginID)
	if err != nil {
		http_server.ResponseError(w, http.StatusBadRequest, ErrInvalidDataTransfer)
		return
	}

	destinationAccountID, err := h.accountUseCase.GetById(body.DestinationAccountID)
	if err != nil {
		http_server.ResponseError(w, http.StatusBadRequest, ErrInvalidDataTransfer)
		return
	}

	input := transfer.CreateTransferInput{
		OriginAccount:      originAccount,
		DestinationAccount: destinationAccountID,
		Amount:             body.Amount,
	}

	output, err := h.useCase.Create(input)

	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrUnexpected):
			h.logger.Err(err).
				Int("Status_code", http.StatusInternalServerError).
				Msg("Occurred when trying to create a transfer.")
			http_server.ResponseError(w, http.StatusInternalServerError, err.Error())
		default:
			h.logger.Err(err).
				Int("Status_code", http.StatusBadRequest).
				Msg("Occurred when trying to create a transfer.")
			http_server.ResponseError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	err = h.accountUseCase.UpdateBalance(input.OriginAccount.AccountID, input.DestinationAccount.AccountID, input.Amount)
	if err != nil {
		http_server.ResponseError(w, http.StatusInternalServerError, ErrUnexpected)
		return
	}

	response := CreateTransferResponse{
		TransferID: output.TransferID,
	}
	h.logger.Info().
		Str("TransferID:", response.TransferID).
		Int("Status:", http.StatusCreated).
		Msg("Transfer created with success")

	http_server.ResponseSuccess(w, http.StatusCreated, response)
}
