package authentication

import (
	"encoding/json"
	auth "github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"net/http"
)

var (
	ErrInvalidPayload = "invalid request payload"
	ErrRequiredFields = "invalid fields"
	ErrNotFound       = "not found"
	ErrUnexpected     = "something went wrong"
)

type LoginRequest struct {
	CPF    string `json:"cpf" validate:"required"`
	Secret string `json:"secret" validate:"required"`
}

type CreateTokenResponse struct {
	Token string `json:"Token" validate:"required"`
}

type LoginResponse struct {
	Token string
}

func (h Handler) CreateToken(w http.ResponseWriter, r *http.Request) {
	var body LoginRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		responseError(w, http.StatusBadRequest, ErrInvalidPayload)
		return
	}

	validator := http_server.NewJSONValidator()
	err := validator.Validate(body)
	if err != nil {
		responseError(w, http.StatusBadRequest, ErrRequiredFields)
		return
	}

	defer r.Body.Close()

	account, err := h.accountUseCase.GetByCpf(body.CPF)
	if err != nil {
		responseError(w, http.StatusNotFound, ErrNotFound)
		return
	}

	input := auth.LoginInputs{
		CPF:     body.CPF,
		Secret:  body.Secret,
		Account: account,
	}

	isValidUser, err := h.useCase.CheckLogin(input)
	if !isValidUser {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	output, err := h.useCase.CreateToken(input)
	if err != nil {
		responseError(w, http.StatusInternalServerError, ErrUnexpected)
		return
	}

	response := CreateTokenResponse{
		Token: output,
	}

	responseSuccess(w, http.StatusCreated, response)
}

func responseError(w http.ResponseWriter, code int, message string) {
	responseSuccess(w, code, map[string]string{"error": message})
}

func responseSuccess(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
