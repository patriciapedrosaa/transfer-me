package account

import (
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
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
		h.logger.Err(err).
			Int("status_code", http.StatusInternalServerError).
			Msg("error occurred when try get accounts")
		http_server.ResponseError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	response := make([]GetAccountResponse, len(accountsList))
	for index, account := range accountsList {
		response[index].AccountID = account.AccountID
		response[index].Name = account.Name
		response[index].CPF = string(account.CPF)
	}
	h.logger.Info().
		Int("status:", http.StatusOK).
		Int("total accounts listed", len(accountsList)).
		Msg("accounts were successfully listed.")

	http_server.ResponseSuccess(w, http.StatusOK, response)
}
