package account

import (
	"errors"
	"github.com/gorilla/mux"
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
			responseError(w, http.StatusNotFound, "not found")
		default:
			responseError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	response := GetBalanceResponse{
		Balance: balance,
	}
	responseSuccess(w, http.StatusOK, response)
}
