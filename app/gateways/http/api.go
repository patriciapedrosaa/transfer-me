package http

import (
	"github.com/gorilla/mux"
	"github.com/patriciapedrosaa/transfer-me/app/gateways/http/account"
	"log"
	"net/http"
)

type Api struct {
	Account account.Handler
}

func NewApi(account account.Handler) Api {
	return Api{
		Account: account,
	}
}

func (a Api) Start(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/accounts/{id}/balance", a.Account.GetBalance).Methods(http.MethodGet)
	r.HandleFunc("/accounts", a.Account.Get).Methods(http.MethodGet)
	r.HandleFunc("/accounts", a.Account.Create).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(port, r))
}
