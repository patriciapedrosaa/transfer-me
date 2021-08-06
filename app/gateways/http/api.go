package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type AccountHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, _ *http.Request)
	GetBalance(w http.ResponseWriter, r *http.Request)
}

type Api struct {
	Account AccountHandler
}

func NewApi(account AccountHandler) Api {
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
