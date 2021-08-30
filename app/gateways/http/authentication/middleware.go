package authentication

import (
	"context"
	"errors"
	http_server "github.com/patriciapedrosaa/transfer-me/app/gateways/http"
	"net/http"
	"strings"
)

var (
	ErrEmptyHeader       = "empty authorization header"
	ErrEmptyToken        = "empty token"
	ErrInvalidAuthMethod = "invalid auth method"
	ErrInvalidToken      = "invalid token"
)

func (h Handler) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.logger.Info().Msg("checking header.")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			h.logger.Err(errors.New(ErrEmptyHeader)).
				Int("status_code", http.StatusBadRequest).
				Msg("error occurred when was validating header")
			http_server.ResponseError(w, http.StatusBadRequest, ErrEmptyHeader)
			return
		}
		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			h.logger.Err(errors.New(ErrEmptyToken)).
				Int("status_code", http.StatusBadRequest).
				Msg("error occurred when was validating token")
			http_server.ResponseError(w, http.StatusBadRequest, ErrEmptyToken)
			return
		}
		if tokenString[0] != "Bearer" {
			h.logger.Err(errors.New(ErrInvalidAuthMethod)).
				Int("status_code", http.StatusBadRequest).
				Msg("error occurred when was validating authentication method")
			http_server.ResponseError(w, http.StatusBadRequest, ErrInvalidAuthMethod)
			return
		}
		validToken, err := h.useCase.ValidatesToken(tokenString[1])
		if err != nil {
			h.logger.Err(errors.New(ErrEmptyHeader)).
				Int("status_code", http.StatusForbidden).
				Msg("error occurred when was validating token")
			http_server.ResponseError(w, http.StatusForbidden, ErrInvalidToken)
			return
		}
		accountID := validToken.Subject

		ctx := context.WithValue(r.Context(), http_server.AccountID, accountID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
