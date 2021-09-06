package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
)

const (
	JsonContentType = "application/json"
)

type RequestContextKey string

var AccountID = RequestContextKey("account_id")

type AccountHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetBalance(w http.ResponseWriter, r *http.Request)
}

type AuthenticationHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Authenticate(next http.HandlerFunc) http.HandlerFunc
}

type TransferHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type Api struct {
	Account  AccountHandler
	Auth     AuthenticationHandler
	Transfer TransferHandler
	Logger   zerolog.Logger
}

func NewApi(account AccountHandler, auth AuthenticationHandler, transfer TransferHandler, logger zerolog.Logger) Api {
	return Api{
		Account:  account,
		Auth:     auth,
		Transfer: transfer,
		Logger:   logger,
	}
}

func (a Api) Start(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/accounts/{id}/balance", a.Account.GetBalance).Methods(http.MethodGet)
	r.HandleFunc("/accounts", a.Account.Get).Methods(http.MethodGet)
	r.HandleFunc("/accounts", a.Account.Create).Methods(http.MethodPost)
	r.HandleFunc("/login", a.Auth.Login).Methods(http.MethodPost)
	r.HandleFunc("/transfers", a.Auth.Authenticate(a.Transfer.Create)).Methods(http.MethodPost)
	r.HandleFunc("/transfers", a.Auth.Authenticate(a.Transfer.Get)).Methods(http.MethodGet)

	if err := http.ListenAndServe(port, r); err != nil {
		a.Logger.Fatal().Err(err).Msg("Startup failed")
	}
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
