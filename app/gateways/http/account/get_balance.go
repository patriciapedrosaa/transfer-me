package account

import (
	"errors"
	"github.com/gorilla/mux"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"net/http"
)

var errNotFound = errors.New("not found")

type GetBalanceResponse struct {
	Balance int `json:"balance"`
}

func (h Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	balance, err := h.useCase.GetBalance(id)
	if err != nil {
		switch err {
		case errNotFound:
			h.logger.Err(err).
				Int("Status_code", http.StatusNotFound).
				Msg("occurred when try get balance")
			http_server.ResponseError(w, http.StatusNotFound, "not found")
		default:
			h.logger.Err(err).
				Int("Status_code", http.StatusInternalServerError).
				Msg("occurred when try get balance")
			http_server.ResponseError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	response := GetBalanceResponse{
		Balance: balance,
	}
	h.logger.Info().
		Int("Status_code", http.StatusOK).
		Msg("Balance was obtained successfully")
	http_server.ResponseSuccess(w, http.StatusOK, response)
}
