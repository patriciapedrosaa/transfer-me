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

type AuthenticationHandler interface {
	CreateToken(w http.ResponseWriter, r *http.Request)
}

type Api struct {
	Account AccountHandler
	Auth    AuthenticationHandler
}

func NewApi(account AccountHandler, auth AuthenticationHandler) Api {
	return Api{
		Account: account,
		Auth:    auth,
	}
}

func (a Api) Start(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/accounts/{id}/balance", a.Account.GetBalance).Methods(http.MethodGet)
	r.HandleFunc("/accounts", a.Account.Get).Methods(http.MethodGet)
	r.HandleFunc("/accounts", a.Account.Create).Methods(http.MethodPost)
	r.HandleFunc("/login", a.Auth.CreateToken).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(port, r))
}
