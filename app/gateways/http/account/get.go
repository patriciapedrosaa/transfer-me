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
			Int("Status_code", http.StatusInternalServerError).
			Msg("occurred when try get accounts")
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
		Int("Status:", http.StatusOK).
		Msgf("Accounts were successfully listed. Total accounts listed: %d", len(accountsList))

	http_server.ResponseSuccess(w, http.StatusOK, response)
}
