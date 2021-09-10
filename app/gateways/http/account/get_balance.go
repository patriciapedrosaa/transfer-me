package account

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account/usecase"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"net/http"
)

var errNotFound = errors.New("not found")

type GetBalanceResponse struct {
	Balance int `json:"balance"`
}

func (h Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	balance, err := h.useCase.GetBalance(ctx, id)
	if err != nil {
		switch err {
		case errNotFound:
			h.logger.Err(err).
				Int("status_code", http.StatusNotFound).
				Msg("error occurred when try get balance")
			http_server.ResponseError(w, http.StatusNotFound, "not found")
		case usecase.ErrInvalidId:
			h.logger.Err(err).
				Int("status_code", http.StatusBadRequest).
				Msg("error occurred when try get balance")
			http_server.ResponseError(w, http.StatusBadRequest, "invalid id format")
		default:
			h.logger.Err(err).
				Int("status_code", http.StatusInternalServerError).
				Msg("error occurred when try get balance")
			http_server.ResponseError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	response := GetBalanceResponse{
		Balance: balance,
	}
	h.logger.Info().
		Int("status_code", http.StatusOK).
		Msg("balance was obtained successfully!")
	http_server.ResponseSuccess(w, http.StatusOK, response)
}
