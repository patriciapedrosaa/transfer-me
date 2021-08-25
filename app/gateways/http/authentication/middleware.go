package authentication

import (
	"context"
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
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http_server.ResponseError(w, http.StatusUnauthorized, ErrEmptyHeader)
			return
		}
		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			http_server.ResponseError(w, http.StatusUnauthorized, ErrEmptyToken)
			return
		}
		if tokenString[0] != "Bearer" {
			http_server.ResponseError(w, http.StatusUnauthorized, ErrInvalidAuthMethod)
			return
		}
		validToken, err := h.useCase.ValidatesToken(tokenString[1])
		if err != nil {
			http_server.ResponseError(w, http.StatusUnauthorized, ErrInvalidToken)
			return
		}
		accountID := validToken.Subject

		ctx := context.WithValue(r.Context(), http_server.AccountID, accountID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
