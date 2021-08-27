package authentication

import (
	"encoding/json"
	"errors"
	auth "github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication/usecase"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"net/http"
)

var (
	ErrInvalidPayload     = "invalid request payload"
	ErrRequiredFields     = "invalid fields"
	ErrInvalidCredentials = "incorrect username or password"
)

type LoginRequest struct {
	CPF    string `json:"cpf" validate:"required"`
	Secret string `json:"secret" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"Token" validate:"required"`
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var body LoginRequest
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
	h.logger.Info().Msgf("Login request: %v", body)

	if err != nil {
		h.logger.Err(err).
			Int("Status_code", http.StatusBadRequest).
			Msg("Occurred when was validating body")
		http_server.ResponseError(w, http.StatusBadRequest, ErrRequiredFields)
		return
	}

	defer r.Body.Close()

	account, err := h.accountUseCase.GetByCpf(body.CPF)
	if err != nil {
		h.logger.Err(err).
			Int("Status_code", http.StatusBadRequest).
			Msg("Occurred when was looking for account")
		http_server.ResponseError(w, http.StatusBadRequest, ErrInvalidCredentials)
		return
	}

	input := auth.LoginInputs{
		CPF:     body.CPF,
		Secret:  body.Secret,
		Account: account,
	}

	output, err := h.useCase.CreateToken(input)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrInvalidCPF):
			h.logger.Err(err).
				Int("Status_code", http.StatusBadRequest).
				Msg("Occurred when was creating token")
			http_server.ResponseError(w, http.StatusBadRequest, ErrInvalidCredentials)
		case errors.Is(err, usecase.ErrInvalidSecret):
			h.logger.Err(err).
				Int("Status_code", http.StatusBadRequest).
				Msg("Occurred when was validating secret")
			http_server.ResponseError(w, http.StatusBadRequest, ErrInvalidCredentials)
		default:
			h.logger.Err(err).
				Int("Status_code", http.StatusInternalServerError).
				Msg("Occurred when was creating token")
			http_server.ResponseError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response := LoginResponse{
		Token: output,
	}
	h.logger.Info().
		Int("Status:", http.StatusCreated).
		Msg("Token created with success")

	http_server.ResponseSuccess(w, http.StatusCreated, response)
}
