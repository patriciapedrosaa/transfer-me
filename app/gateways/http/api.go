package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	JsonContentType = "application/json"
)

type AccountHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, _ *http.Request)
	GetBalance(w http.ResponseWriter, r *http.Request)
}

type AuthenticationHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
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
	r.HandleFunc("/login", a.Auth.Login).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(port, r))
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseSuccess(w, code, map[string]string{"error": message})
}

func ResponseSuccess(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
