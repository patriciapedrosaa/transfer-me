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

	account, err := h.accountUseCase.GetByCpf(body.CPF)
	if err != nil {
		http_server.ResponseError(w, http.StatusBadRequest, usecase.ErrInvalidCredentials.Error())
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
		case errors.Is(err, usecase.ErrInvalidCredentials):
			http_server.ResponseError(w, http.StatusBadRequest, err.Error())
		default:
			http_server.ResponseError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response := LoginResponse{
		Token: output,
	}

	http_server.ResponseSuccess(w, http.StatusCreated, response)
}
