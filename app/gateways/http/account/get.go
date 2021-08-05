package account

import (
	"net/http"
)

type GetAccountResponse struct {
	AccountID string `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
}

func (h Handler) Get(w http.ResponseWriter, _ *http.Request) {
	accountsList, err := h.useCase.GetAccounts()
	if err != nil {
		responseError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	response := make([]GetAccountResponse, len(accountsList))
	for index, account := range accountsList {
		response[index].AccountID = account.AccountID
		response[index].Name = account.Name
		response[index].CPF = string(account.CPF)
	}

	responseSuccess(w, http.StatusOK, response)
}
