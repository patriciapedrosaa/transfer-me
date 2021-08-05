package account

import (
	"encoding/json"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"net/http"
)

var ErrInvalidPayload = "invalid request payload"

type CreateAccountRequest struct {
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type CreateAccountResponse struct {
	AccountID string `json:"id"`
	Name      string `json:"name"`
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body CreateAccountRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&body); err != nil {
		responseError(w, http.StatusBadRequest, ErrInvalidPayload)
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
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := CreateAccountResponse{
		AccountID: output.AccountID,
		Name:      output.Name,
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
