package account

import (
	"encoding/json"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"net/http"
)

var (
	ErrInvalidPayload = "invalid request payload"
	ErrRequiredFields = "invalid fields"
)

type CreateAccountRequest struct {
	Name   string `json:"name" validate:"required"`
	CPF    string `json:"cpf" validate:"required"`
	Secret string `json:"secret" validate:"required"`
}

type CreateAccountResponse struct {
	AccountID string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body CreateAccountRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		h.logger.Err(err).
			Int("Status_code", http.StatusBadRequest).
			Msg("Occurred when decoding body")
		http_server.ResponseError(w, http.StatusBadRequest, ErrInvalidPayload)
		return
	}

	validator := http_server.NewJSONValidator()
	err := validator.Validate(body)
	h.logger.Info().Msgf("Create account request: %v", body)
	if err != nil {
		h.logger.Err(err).
			Int("Status_code", http.StatusBadRequest).
			Msg("Occurred when was validating body")
		http_server.ResponseError(w, http.StatusBadRequest, ErrRequiredFields)
		return
	}

	defer r.Body.Close()

	input := account.CreateAccountInput{
		Name:   body.Name,
		Secret: body.Secret,
		CPF:    body.CPF,
	}

	output, err := h.useCase.Create(input)
	if err != nil {
		h.logger.Err(err).
			Int("Status_code", http.StatusBadRequest).
			Msg("Occurred when trying to create an account.")
		http_server.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := CreateAccountResponse{
		AccountID: output.AccountID,
		Name:      output.Name,
	}
	h.logger.Info().
		Str("AccountID:", response.AccountID).
		Int("Status:", http.StatusCreated).
		Msg("Account created with success")

	http_server.ResponseSuccess(w, http.StatusCreated, response)
}
